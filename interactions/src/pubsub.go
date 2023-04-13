package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

type VideoNotification struct {
	VideoURL         string
	InteractionToken string
}

var Client *pubsub.Client

func init() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "peter-built")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Client = client
}

func publishVideoNotification(notification *VideoNotification) {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "peter-built")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	topicID := "webm-videos"

	data, err := json.Marshal(notification)
	if err != nil {
		log.Printf("couldn't unmarshal %v", err)
		return
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	_, err = result.Get(ctx)
	if err != nil {
		fmt.Printf("couldn't publish: %v", err)
	}
}
