package main

import (
	"fmt"
	"os"
	"time"

	"github.com/programminh/wright/qpx"
)

type Trip interface {
	Price() float64
	Stops() int
	Duration() time.Duration
	Carrier() string
}

func main() {
	var (
		trip Trip
		err  error
	)

	fmt.Println("Finding cheapest flight from EDI to YUL on 2017-05-19")
	if trip, err = qpx.Cheapest("EDI", "YUL", "2017-05-19"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Found %.2f$ - takes %s with %d planes by %s\n", trip.Price(), trip.Duration(), trip.Stops(), trip.Carrier())

	fmt.Println("Finding cheapest flight from GLA to YUL on 2017-05-19")
	if trip, err = qpx.Cheapest("GLA", "YUL", "2017-05-19"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Found %.2f$ - takes %s with %d planes by %s\n", trip.Price(), trip.Duration(), trip.Stops(), trip.Carrier())
}
