package main

import (
	"log"
	"net/http"

	"github.com/saurzv/visitc/pkg/analytics"
	"github.com/saurzv/visitc/pkg/http/rest"
	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/storage/json"
)

func main() {
	s, err := json.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	lister := listing.NewService(s)
	analyzer := analytics.NewService(s)

	r := rest.Handler(lister, analyzer)

	http.ListenAndServe(":8080", r)
}
