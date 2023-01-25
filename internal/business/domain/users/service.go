package users

type userProvider interface {
	FindAll() ([]User, error)
}

type Service struct {
	provider userProvider
}

func NewService(provider userProvider) *Service {
	return &Service{provider: provider}
}

func (s *Service) IsUserExists(login string) (bool, error) {
	users, err := s.provider.FindAll()
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.Login == login {
			return true, nil
		}
	}

	return false, nil
}
