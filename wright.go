package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/programminh/wright/qpx"
)

type Trip interface {
	Price() float64
	Stops() int
	Duration() time.Duration
	Carrier() string
}

var green = "#66BB6A"
var darkgreen = "#1B5E20"

func main() {
	var (
		edi, gla Trip
		err      error
		date     = "2017-05-19"
		dst      = "YUL"
	)

	if edi, err = qpx.Cheapest("EDI", dst, date); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	slack("EDI", dst, date, green, edi)

	if gla, err = qpx.Cheapest("GLA", dst, date); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	slack("GLA", dst, date, darkgreen, gla)
}

type message struct {
	Username    string       `json:"username"`
	Channel     string       `json:"channel"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []attachment `json:"attachments"`
}

type attachment struct {
	Fallback   string   `json:"fallback"`
	Title      string   `json:"title"`
	Text       string   `json:"text"`
	Color      string   `json:"color"`
	Carrier    string   `json:"carrier"`
	Footer     string   `json:"footer"`
	MarkdownIn []string `json:"mrkdwn_in"`
}

func slack(origin, destination, date, color string, trip Trip) {
	msg := message{
		Username:  "Wright",
		Channel:   "#general",
		IconEmoji: ":google:",
	}

	att := attachment{
		Color:    color,
		Fallback: fmt.Sprintf("Cheapest flight from %s to %s on %s", origin, destination, date),
		Title:    fmt.Sprintf("%s to %s on %s - %.2f$", origin, destination, date, trip.Price()),
		Text:     fmt.Sprintf("Duration %s with %d stops", strings.TrimSuffix(fmt.Sprintf("%s", trip.Duration()), "m0s"), trip.Stops()),
		Footer:   fmt.Sprintf("by %s", trip.Carrier()),
	}

	msg.Attachments = []attachment{att}

	buf := bytes.NewBuffer(nil)

	json.NewEncoder(buf).Encode(msg)

	if _, err := http.Post(endpoint, "application/json", buf); err != nil {
		log.Println(err)
	}
}
