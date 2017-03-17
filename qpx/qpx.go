package qpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/programminh/wright/clerk"
)

const provider = "QPX"

func Cheapest(origin, destination, date string) (to tripOption, err error) {
	var (
		res  *http.Response
		body response
		buf  = bytes.NewBuffer(nil)
		b    []byte
	)

	payload := newRequest(origin, destination, date)

	if b, err = json.MarshalIndent(payload, "", " "); err != nil {
		return
	}

	if res, err = http.Post(endpoint, "application/json", bytes.NewBuffer(b)); err != nil {
		return
	}

	tee := io.TeeReader(res.Body, buf)

	if err = json.NewDecoder(tee).Decode(&body); err != nil {
		return
	}
	res.Body.Close()

	body.sort()
	to = body.Cheapest()

	t, _ := time.Parse("2006-01-02", date)

	clerk.Log(provider, origin, destination, string(b), buf.String(), t)

	return
}
