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
	var dst = "YUL"

	search("EDI", dst, "2017-05-19", green)
	search("EDI", dst, "2017-05-20", green)

	search("GLA", dst, "2017-05-19", darkgreen)
	search("GLA", dst, "2017-05-20", darkgreen)
}

func search(origin, dst, date, color string) {
	var (
		trip Trip
		err  error
	)

	if trip, err = qpx.Cheapest(origin, dst, date); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	slack(origin, dst, date, color, trip)
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
