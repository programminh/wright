// +build darwin windows
package qpx

import (
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
