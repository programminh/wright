package qpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func logfile() (f *os.File, err error) {
	filename := time.Now().Format("2006-01-02_1504.log")

	path := filepath.Join(filepath.Join(os.Getenv("HOME"), "Desktop", "wright", "qpx", filename))
	dir := filepath.Dir(filename)

	if _, err = os.Stat(dir); err != nil {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return
		}
	}

	return os.Create(path)
}

func Cheapest(origin, destination, date string) (to tripOption, err error) {
	var (
		res  *http.Response
		f    *os.File
		resp response
		buf  = bytes.NewBuffer(nil)
	)

	payload := newRequest(origin, destination, date)

	if f, err = logfile(); err != nil {
		return
	}
	defer f.Close()

	w := io.MultiWriter(buf, f)

	if err = json.NewEncoder(w).Encode(payload); err != nil {
		return
	}

	if res, err = http.Post(endpoint, "application/json", buf); err != nil {
		return
	}

	f.WriteString("===========================\n")

	tee := io.TeeReader(res.Body, f)

	if err = json.NewDecoder(tee).Decode(&resp); err != nil {
		return
	}
	res.Body.Close()

	resp.sort()
	to = resp.Cheapest()

	return
}
