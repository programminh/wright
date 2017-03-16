package qpx

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

type carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type responseSlice struct {
	Duration int       `json:"duration"`
	Segment  []segment `json:"segment"`
}

type segment struct {
	Flight struct {
		Carrier string `json:"carrier"`
	} `json:"flight"`
}

type tripOption struct {
	CarrierName string
	ID          string          `json:"id"`
	SaleTotal   string          `json:"saleTotal"`
	Slice       []responseSlice `json:"slice"`
}

func (to tripOption) Carrier() string {
	return to.CarrierName
}

func (to tripOption) Stops() int {
	return len(to.Slice[0].Segment) - 1
}

func (to tripOption) Price() (price float64) {
	var err error

	if price, err = strconv.ParseFloat(to.SaleTotal[3:], 64); err != nil {
		log.Printf("can't parse amount (%s - %s): (%s)", to.SaleTotal, to.SaleTotal[4:], err)
	}

	return
}

func (to tripOption) Duration() time.Duration {
	var (
		d   time.Duration
		err error
	)

	if d, err = time.ParseDuration(fmt.Sprintf("%dm", to.Slice[0].Duration)); err != nil {
		log.Printf("can't parse duration %d", to.Slice[0].Duration)
	}

	return d
}

type tripOptions []tripOption

func (to tripOptions) Len() int           { return len(to) }
func (to tripOptions) Swap(i, j int)      { to[i], to[j] = to[j], to[i] }
func (to tripOptions) Less(i, j int) bool { return to[i].Price() < to[j].Price() }

type response struct {
	Trips struct {
		Data struct {
			Carrier []carrier `json:"carrier"`
		} `json:"data"`
		TripOption tripOptions `json:"tripOption"`
	} `json:"trips"`

	sorted bool
}

func (r *response) sort() {
	if r.sorted {
		return
	}

	r.sorted = true

	sort.Sort(r.Trips.TripOption)
}

func (r *response) Cheapest() tripOption {
	// Assume the carrier name is the first flight
	to := r.Trips.TripOption[0]
	name := to.Slice[0].Segment[0].Flight.Carrier

	to.CarrierName = r.Carrier(name)

	return to
}

func (r *response) Prices() []float64 {
	arr := make([]float64, len(r.Trips.TripOption))

	for i, trip := range r.Trips.TripOption {
		arr[i] = trip.Price()
	}

	return arr
}

func (r *response) Carrier(s string) string {
	for _, c := range r.Trips.Data.Carrier {
		if s == c.Code {
			return c.Name
		}
	}

	return ""
}
