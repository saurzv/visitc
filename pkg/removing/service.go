package removing

type Repository interface {
	RemoveSite(string) error
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
	if err := s.r.RemoveSite(id); err != nil {
		return err
	}
	return nil
}
