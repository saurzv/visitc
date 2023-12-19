package main

import (
	"fmt"
	"log"

	"github.com/saurzv/visitc/pkg/counting"
	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/storage/json"
)

func main() {
	s, err := json.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	// adder := adding.NewService(s)
	lister := listing.NewService(s)
	// remover := removing.NewService(s)
	counter := counting.NewService(s)

	counter.IncreaseCount("Github_b75dd37a6648af45")

	// err = adder.AddSite(adding.DefaultSite...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lister.GetSite("Github_b75dd37a6648af45"))
	// fmt.Println(lister.GetSites())
}
