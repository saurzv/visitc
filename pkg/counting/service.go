package counting

type Repository interface {
	IncreaseCount(string) error
}

type Service interface {
	IncreaseCount(string) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) IncreaseCount(id string) error {
	if err := s.r.IncreaseCount(id); err != nil {
		return err
	}
	return nil
}
