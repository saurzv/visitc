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

	if err := adder.AddSite(adding.DefaultSite); err != nil {
		log.Fatal(err)
	}
	// if err := adder.RemoveSite("Github-2_0c046b55ff94e30e"); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(lister.GetSites())
}
