package qpx

import (
	"io"
	"os"
)

func logfile() (wc io.WriteCloser, err error) {
	return os.Stdout, nil
}
