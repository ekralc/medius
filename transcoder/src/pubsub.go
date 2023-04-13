package main

import (
	"encoding/json"
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

func PubsubPush(c *gin.Context) {
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

	if err := HandleVideoNotification(&notification); err != nil {
		log.Printf("error occurred during video transcoding: %v", err)
	}

	c.Status(http.StatusAccepted)
}
