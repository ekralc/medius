package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"os/exec"
	"path/filepath"
)

var ApplicationID string = "1079541448820662412"

func transcode(URL string, outputPath string) {
	cmd := exec.Command("ffmpeg", "-y", "-i", URL, "-movflags", "faststart", "-preset", "superfast", outputPath)
	stdout, err := cmd.StdoutPipe()
	stdoutReader := bufio.NewReader(stdout)
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrReader := bufio.NewReader(stderr)

	go handleReader(stdoutReader)
	go handleReader(stderrReader)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func generateVideoName() (string, error) {
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", err
	}

	prefix := "video-"
	hex := hex.EncodeToString(randBytes)
	suffix := ".mp4"

	return filepath.Join(os.TempDir(), prefix+hex+suffix), nil
}

func HandleVideoNotification(notification *NotificationPayload) error {
	url := buildWebhookEndpoint(ApplicationID, notification.InteractionToken)

	client := http.Client{}

	reader, writer := io.Pipe()
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
		return err
	}

	mwriter := multipart.NewWriter(writer)
	req.Header.Set("Content-Type", mwriter.FormDataContentType())

	errchan := make(chan error)

	// start transcoding
	path, _ := generateVideoName()
	transcode(notification.VideoURL, path)

	go func() {
		defer close(errchan)
		defer writer.Close()
		defer mwriter.Close()

		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "files[0]", "video.mp4"))
		h.Set("Content-Type", "video/mp4")

		w, err := mwriter.CreatePart(h)
		if err != nil {
			errchan <- err
			return
		}

		in, err := os.Open(path)
		if err != nil {
			errchan <- err
			return
		}
		defer in.Close()

		if written, err := io.Copy(w, in); err != nil {
			errchan <- fmt.Errorf("error copying %s (%d bytes written): %v", path, written, err)
			return
		}
		if err := mwriter.Close(); err != nil {
			errchan <- err
			return
		}
	}()

	resp, err := client.Do(req)
	merr := <-errchan
	if err != nil || merr != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		return fmt.Errorf("http error uploading video status code %d: %v", resp.StatusCode, string(body))
	}

	return nil
}
