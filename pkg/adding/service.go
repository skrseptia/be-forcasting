package adding

type RepositoryMySQL interface {
	CreateUser(User) (uint, error)
}

type Service interface {
	AddUser(User) (User, error)
}

type service struct {
	rmy RepositoryMySQL
}

func NewService(rmy RepositoryMySQL) Service {
	return &service{rmy}
}

func (s *service) AddUser(au User) (User, error) {
	var err error
	au.ID, err = s.rmy.CreateUser(au)
	if err != nil {
		return au, err
	}

	return au, nil
}
