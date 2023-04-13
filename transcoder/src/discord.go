package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// buildWebhookEndpoint constructs a Discord API endpoint for an interaction
func buildWebhookEndpoint(application string, token string) string {
	url := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s/messages/@original", ApplicationID, token)

	return url
}

func sendFollowupMessage(token string, message string) error {
	url := buildWebhookEndpoint(ApplicationID, token)

	values := map[string]string{
		"content": message,
	}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		return err
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

}
