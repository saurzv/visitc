package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saurzv/visitc/pkg/analytics"
	"github.com/saurzv/visitc/pkg/listing"
	"github.com/saurzv/visitc/pkg/rendering"
)

func Handler(l listing.Service, a analytics.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/api/{id}", getSVG(l, a))

	return router
}

func getSVG(l listing.Service, a analytics.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		siteID := chi.URLParam(r, "id")
		site, err := l.GetSite(siteID)
		if err != nil {
			fmt.Println(err)
			return
		}

		go func() {
			addr := r.RemoteAddr
			err := a.UpdateLastVisit(addr)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = a.IncreaseCount(siteID)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()

		svg, err := rendering.SVG(site)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(svg)
	}
}
