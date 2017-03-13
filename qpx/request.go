package qpx

type Passengers struct {
	Kind              string `json:"kind"`
	AdultCount        int    `json:"adultCount"`
	ChildCount        int    `json:"childCount"`
	InfantInLapCount  int    `json:"infantInLapCount"`
	InfantInSeatCount int    `json:"infantInSeatCount"`
	SeniorCount       int    `json:"seniorCount"`
}

type Slice struct {
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

type Request struct {
	Request struct {
		Passengers       Passengers `json:"passengers"`
		Slice            []Slice    `json:"slice"`
		MaxPrice         string     `json:"maxPrice"`
		SaleCountry      string     `json:"saleCountry"`
		TicketingCountry string     `json:"ticketingCountry"`
		Refundable       string     `json:"refundable"`
		Solutions        int        `json:"solutions"`
	} `json:"request"`
}
