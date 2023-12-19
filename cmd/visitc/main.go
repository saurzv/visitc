package main

import (
	"fmt"
	"log"

	"github.com/saurzv/visitc/pkg/adding"
	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/storage/json"
)

func main() {
	s, err := json.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	adder := adding.NewService(s)
	lister := listing.NewService(s)
	// remover := removing.NewService(s)

	err = adder.AddSite(adding.DefaultSite...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lister.GetSites())
}
