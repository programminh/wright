package qpx

type passengers struct {
	Kind              string `json:"kind"`
	AdultCount        int    `json:"adultCount"`
	ChildCount        int    `json:"childCount"`
	InfantInLapCount  int    `json:"infantInLapCount"`
	InfantInSeatCount int    `json:"infantInSeatCount"`
	SeniorCount       int    `json:"seniorCount"`
}

type requestSlice struct {
	Kind                   string `json:"string"`
	Origin                 string `json:"origin"`
	Destination            string `json:"destination"`
	Date                   string `json:"date"`
	MaxStop                int    `json:"maxStop"`
	MaxConnectionDuration  int    `json:"maxConnectionDuration"`
	PreferredCabin         string `json:"preferredCabin"`
	PermittedDepartureTime struct {
		Kind        string `json:"kind"`
		EarlierTime string `json:"earliestTime"`
		LatestTime  string `json:"latestTime"`
	} `json:"permittedDepartureTime"`
	PermittedCarrier  []string `json:"permittedCarrier"`
	Alliance          string   `json:"alliance"`
	ProhibitedCarrier []string `json:"prohibitedCarrier"`
}

type request struct {
	Request struct {
		Passengers       passengers     `json:"passengers"`
		Slice            []requestSlice `json:"slice"`
		MaxPrice         string         `json:"maxPrice"`
		SaleCountry      string         `json:"saleCountry"`
		TicketingCountry string         `json:"ticketingCountry"`
		Refundable       string         `json:"refundable"`
		Solutions        int            `json:"solutions"`
	} `json:"request"`
}

func newRequest(origin, destination, date string) request {
	req := request{}
	slice := requestSlice{
		Kind:                  "qpxexpress#sliceInput",
		Origin:                origin,
		Destination:           destination,
		Date:                  date,
		MaxStop:               3,
		MaxConnectionDuration: 300,
	}

	pass := passengers{
		Kind:       "qpxexpress#passengerCounts",
		AdultCount: 1,
	}

	req.Request.SaleCountry = "CA"
	req.Request.Solutions = 50
	req.Request.Slice = []requestSlice{slice}
	req.Request.Passengers = pass

	return req
}
