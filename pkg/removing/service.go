package removing

import (
	"github.com/saurzv/visitc/pkg/listing"
)

type Repository interface {
	GetSite(string) (listing.Site, error)
	RemoveSite(listing.Site) error
}

type Service interface {
	RemoveSite(string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) RemoveSite(id string) error {
	availableSite, err := s.r.GetSite(id)
	if err != nil {
		return err
	}
	if err := s.r.RemoveSite(availableSite); err != nil {
		return err
	}
	return nil
}
