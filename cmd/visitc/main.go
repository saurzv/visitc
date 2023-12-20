package main

import (
	"log"

	"github.com/saurzv/visitc/pkg/analytics"
	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/rendering"
	"github.com/saurzv/visitc/pkg/storage/json"
)

func main() {
	s, err := json.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	// updator := updating.NewService(s)
	lister := listing.NewService(s)
	// remover := removing.NewService(s)
	// counter := counting.NewService(s)

	analyzer := analytics.NewService(s)

	err = analyzer.IncreaseCount("blog_ce93042ea4b3a5d9")

	// counter.IncreaseCount("Github_b75dd37a6648af45")

	// err = updator.AddSite(updating.DefaultSite...)
	if err != nil {
		log.Fatal(err)
	}
	site, _ := lister.GetSite("blog_ce93042ea4b3a5d9")
	// fmt.Print(site.Human())
	err = rendering.SVG(site)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(lister.GetSites())
}
