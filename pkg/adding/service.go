package adding

import (
	"errors"
	"fmt"

	"github.com/saurzv/visitc/pkg/listing"
)

type Service interface {
	AddSite(...Site) error
}

type Repository interface {
	AddSite(Site) error
	GetAllSites() []listing.Site
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddSite(sites ...Site) error {
	allSites := s.r.GetAllSites()

	mp := map[string]bool{}
	for _, site := range allSites {
		mp[site.Name] = true
	}

	for _, site := range sites {
		if mp[site.Name] {
			errMsg := fmt.Sprintf("Site '%s' already exists", site.Name)
			return errors.New(errMsg)
		}
	}

	for _, site := range sites {
		s.r.AddSite(site)
	}
	return nil
}
