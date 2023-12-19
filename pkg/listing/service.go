package listing

type Service interface {
	GetSite(string) Site
	GetSites() []Site
}

type Repository interface {
	GetSite(string) Site
	GetAllSites() []Site
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetSites() []Site {
	return s.r.GetAllSites()
}

func (s *service) GetSite(id string) Site {
	return s.r.GetSite(id)
}
