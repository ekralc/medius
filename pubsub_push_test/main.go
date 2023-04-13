package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type NotificationPayload struct {
	VideoURL         string
	InteractionToken string
}

func HelloPubSub(c *gin.Context) {
	var m PubSubMessage
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// byte slice unmarshalling handles base64 decoding
	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var notification NotificationPayload
	if err := json.Unmarshal(m.Message.Data, &notification); err != nil {
		log.Printf("json.Unmarshal %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	applicationID := "560211485687808011"
	url := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s/messages/@original", applicationID, notification.InteractionToken)

	client := http.Client{}

	payload, err := json.Marshal(map[string]interface{}{
		"content": "fucking kill urself",
	})
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

	msg := string(m.Message.Data)
	log.Printf("hello %v", msg)
}

func main() {
	r := gin.Default()
	r.POST("/pubsub", HelloPubSub)
	r.Run()
}
