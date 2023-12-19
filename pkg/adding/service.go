package adding

import (
	"errors"

	"github.com/saurzv/visitc/pkg/listing"
)

var ErrDuplicate = errors.New("site alreay exists")

type Service interface {
	AddSite(Site) error
	RemoveSite(string) error
}

type Repository interface {
	AddSite(Site) error
	RemoveSite(string) error
	GetAllSites() []listing.Site
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddSite(newSite Site) error {
	sites := s.r.GetAllSites()
	for _, site := range sites {
		if site.Name == newSite.Name {
			return ErrDuplicate
		}
	}

	if err := s.r.AddSite(newSite); err != nil {
		return err
	}
	return nil
}

func (s *service) RemoveSite(id string) error {
	return s.r.RemoveSite(id)
}
