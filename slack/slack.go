package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Message struct {
	Username    string       `json:"username"`
	Channel     string       `json:"channel"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback   string   `json:"fallback"`
	Title      string   `json:"title"`
	Text       string   `json:"text"`
	Color      string   `json:"color"`
	Carrier    string   `json:"carrier"`
	Footer     string   `json:"footer"`
	MarkdownIn []string `json:"mrkdwn_in"`
}

func Send(m Message) error {
	buf := bytes.NewBuffer(nil)

	json.NewEncoder(buf).Encode(m)

	_, err := http.Post(endpoint, "application/json", buf)

	return err
}
