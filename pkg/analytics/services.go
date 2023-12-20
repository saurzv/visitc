package analytics

import (
	"errors"
	"fmt"
	"time"
)

type Repository interface {
	GetLastVisit(string) (time.Time, error)
	UpdateLastVisit(string) error
	AddVisit(string) error
	IncreaseCount(string) error
}

type Service interface {
	IncreaseCount(string) error
	isValidVisit(string) error // TO-DO: Need a definition of valid visits
	UpdateLastVisit(string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateLastVisit(addr string) error {
	// if s.isValidVisit(addr) {
	// 	s.r.UpdateLastVisit(addr)
	// }
	err := s.r.AddVisit(addr)
	if err != nil {
		fmt.Println(err)
	}

	if err = s.isValidVisit(addr); err != nil {
		return err
	}

	s.r.UpdateLastVisit(addr)
	return nil
}

func (s *service) isValidVisit(ip string) error {
	t, err := s.r.GetLastVisit(ip)
	if err != nil {
		return err
	}

	if time.Since(t).Hours() > 1 {
		return nil
	}

	return errors.New("invalid visit")
}

func (s *service) IncreaseCount(id string) error {
	if err := s.r.IncreaseCount(id); err != nil {
		return err
	}

	return nil
}
