package main

import (
	"time"

	"github.com/programminh/wright"
)

func main() {
	t := wright.Trip{
		AdultCount:  2,
		Origin:      "EDI",
		Destination: "YUL",
		Date:        time.Date(2017, time.Month(5), 19, 0, 0, 0, 0, time.UTC),
	}

	wright.Search(&t)
}
