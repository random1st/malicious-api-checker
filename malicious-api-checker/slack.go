package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const WEBHOOKURL = "SLACK_WEBHOOK_URL"

var (
	ErrInvalidStatusCode    = errors.New("invalid status code")
	ErrSlackWebhookNotFound = errors.New("slack webhook not found in env variables")
)

type SlackPayload struct {
	Text      string `json:"text"` // To create a link in your text, enclose the URL in <> angle brackets
	Username  string `json:"username,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Channel   string `json:"channel,omitempty"`
}

type slackAPI struct {
}

func createSlackAPI() *slackAPI {
	return &slackAPI{}
}

func (*slackAPI) sendMessage(subject string, message string) error {
	slackURL, found := os.LookupEnv(WEBHOOKURL)
	if !found {
		log.Println("No Slack webhook url provided")
		return ErrSlackWebhookNotFound
	}
	payload := SlackPayload{
		Text: fmt.Sprintf("%s: %s", subject, message),
	}
	payloadJSON, _ := json.Marshal(payload)
	resp, err := http.Post(slackURL, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Println("Can't send Slack message")
		return err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Invalid status code")
		return ErrInvalidStatusCode
	}
	return nil

}
