package main

import (
	"fmt"
	"log"
	"os"

	"github.com/programminh/wright"
	"github.com/programminh/wright/qpx"
)

func main() {
	var (
		res qpx.Response
		err error
	)

	if res, err = wright.Search("EDI", "YUL", "2017-05-19"); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Found %d results\n", len(res.Trips.TripOption))
	wright.PrettyPrint(res.Cheapest())
	wright.PrettyPrint(res.Mean())
	fmt.Println(res.Prices())

}
