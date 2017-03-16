package qpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Cheapest(origin, destination, date string) (to tripOption, err error) {
	var (
		res  *http.Response
		f    io.WriteCloser
		resp response
		buf  = bytes.NewBuffer(nil)
	)

	payload := newRequest(origin, destination, date)

	if f, err = logfile(); err != nil {
		return
	}

	defer func() {
		if f != os.Stdout {
			f.Close()
		}
	}()

	w := io.MultiWriter(buf, f)

	if err = json.NewEncoder(w).Encode(payload); err != nil {
		return
	}

	if res, err = http.Post(endpoint, "application/json", buf); err != nil {
		return
	}

	f.Write([]byte("===========================\n"))

	tee := io.TeeReader(res.Body, f)

	if err = json.NewDecoder(tee).Decode(&resp); err != nil {
		return
	}
	res.Body.Close()

	resp.sort()
	to = resp.Cheapest()

	return
}
