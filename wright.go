package wright

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/programminh/wright/qpx"
)

var endpoint = fmt.Sprintf("https://www.googleapis.com/qpxExpress/v1/trips/search?key=%s", APIKey)

func Search(origin, destination, date string) (resp qpx.Response, err error) {
	var (
		b        []byte
		res      *http.Response
		f        *os.File
		filename = time.Now().Format("2006-01-02_1504.log")
	)

	payload := qpx.NewRequest(origin, destination, date)

	if f, err = os.Create(filepath.Join(os.Getenv("HOME"), "Desktop/wright", filename)); err != nil {
		return
	}

	if b, err = json.Marshal(payload); err != nil {
		return
	}

	PrettyPrint(payload)

	if res, err = http.Post(endpoint, "application/json", bytes.NewBuffer(b)); err != nil {
		return
	}

	tee := io.TeeReader(res.Body, f)

	if err = json.NewDecoder(tee).Decode(&resp); err != nil {
		return
	}
	res.Body.Close()

	return
}

func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(b))
}
