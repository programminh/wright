package clerk

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

var client *datastore.Client
var ctx = context.Background()

func init() {
	var err error

	if client, err = datastore.NewClient(ctx, "wright-154519"); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

type search struct {
	Provider    string
	Origin      string
	Destination string
	Date        time.Time
	Request     string `datastore:",noindex"`
	Response    string `datastore:",noindex"`
	Created     time.Time
}

func Log(provider, origin, destination, req, res string, date time.Time) {
	var err error

	s := search{
		Provider:    provider,
		Origin:      origin,
		Destination: destination,
		Date:        date,
		Request:     req,
		Response:    res,
		Created:     time.Now(),
	}

	key := datastore.IncompleteKey("Searches", nil)

	if _, err = client.Put(ctx, key, &s); err != nil {
		log.Println(err)
	}
}
