package user

type Service interface {
	FindByName(username string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByName(username string) (User, error) {
	user, err := s.repository.FindByName(username)
	return user, err
}
