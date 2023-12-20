package analytics

type Repository interface {
	IncreaseCount(string) error
}

type Service interface {
	IncreaseCount(string) error
	IsValidVisit(string) bool // TO-DO: Need a definition of valid visits
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) IsValidVisit(ip string) bool {
	return false
}

func (s *service) IncreaseCount(id string) error {
	if err := s.r.IncreaseCount(id); err != nil {
		return err
	}
	return nil
}
