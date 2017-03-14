package qpx

import (
	"log"
	"os"
	"sort"
	"strconv"
)

type Carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ResponseSlice struct {
	ID       string    `json:"id"`
	Duration int       `json:"duration"`
	Segment  []Segment `json:"segment"`
}

type Segment struct {
	Flight struct {
		Carrier string `json:"carrier"`
	} `json:"flight"`
}

type TripOption struct {
	SaleTotal string          `json:"saleTotal"`
	Slice     []ResponseSlice `json:"slice"`
}

func (to TripOption) Price() (price float64) {
	var err error

	if price, err = strconv.ParseFloat(to.SaleTotal[3:], 64); err != nil {
		log.Println("can't parse amount (%s - %s): (%s)", to.SaleTotal, to.SaleTotal[4:], err)
		os.Exit(1)
	}

	return
}

type TripOptions []TripOption

func (to TripOptions) Len() int           { return len(to) }
func (to TripOptions) Swap(i, j int)      { to[i], to[j] = to[j], to[i] }
func (to TripOptions) Less(i, j int) bool { return to[i].Price() < to[j].Price() }

type Response struct {
	Trips struct {
		Carrier    []Carrier   `json:"carrier"`
		TripOption TripOptions `json:"tripOption"`
	} `json:"trips"`

	sorted bool
}

func (r *Response) sort() {
	if r.sorted {
		return
	}

	r.sorted = true

	sort.Sort(r.Trips.TripOption)
}

func (r *Response) Cheapest() TripOption {
	r.sort()

	return r.Trips.TripOption[0]
}

func (r *Response) Prices() []float64 {
	arr := make([]float64, len(r.Trips.TripOption))

	for i, trip := range r.Trips.TripOption {
		arr[i] = trip.Price()
	}

	return arr
}

func (r *Response) Mean() TripOption {
	r.sort()

	return r.Trips.TripOption[len(r.Trips.TripOption)/2]
}
