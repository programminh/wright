package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/programminh/wright/qpx"
	"github.com/programminh/wright/slack"
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
	echo(origin, dst, date, color, trip)
}

func echo(origin, destination, date, color string, trip Trip) {
	msg := slack.Message{
		Username:  "Wright",
		Channel:   "#general",
		IconEmoji: ":google:",
	}

	att := slack.Attachment{
		Color:    color,
		Fallback: fmt.Sprintf("Cheapest flight from %s to %s on %s", origin, destination, date),
		Title:    fmt.Sprintf("%s to %s - %.2f$", origin, destination, trip.Price()),
		Text:     fmt.Sprintf("Duration %s with %d stops on %s", strings.TrimSuffix(fmt.Sprintf("%s", trip.Duration()), "m0s"), trip.Stops(), date),
		Footer:   fmt.Sprintf("by %s", trip.Carrier()),
	}

	msg.Attachments = []slack.Attachment{att}
}
