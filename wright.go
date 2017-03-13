package wright

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/programminh/wright/qpx"
)

var endpoint = fmt.Sprintf("https://www.googleapis.com/qpxExpress/v1/trips/search?key=%", APIKey)

type Trip struct {
	AdultCount  int
	Origin      string
	Destination string
	Date        time.Time
}

func (t Trip) MarshalJSON() (b []byte, err error) {
	req := qpx.Request{}
	slice := qpx.Slice{
		Kind:                  "qpxexpress#sliceInput",
		Origin:                t.Origin,
		Destination:           t.Destination,
		Date:                  t.Date.Format("2006-01-02"),
		MaxStop:               3,
		MaxConnectionDuration: 300,
	}

	pass := qpx.Passengers{
		Kind:       "qpxexpress#passengerCounts",
		AdultCount: t.AdultCount,
	}

	req.Request.Solutions = 10
	req.Request.Slice = []qpx.Slice{slice}
	req.Request.Passengers = pass

	b, err = json.Marshal(req)

	return
}

func Search(t *Trip) (err error) {
	var (
		b   []byte
		res *http.Response
	)

	if b, err = json.Marshal(t); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prettyprint(t)

	if res, err = http.Post(endpoint, "application/json", bytes.NewBuffer(b)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, res.Body)
	res.Body.Close()

	return
}

func prettyprint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(b))
}
