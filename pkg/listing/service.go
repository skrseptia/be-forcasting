package listing

type RepositoryMySQL interface {
}

type Service interface {
}

type service struct {
	rmy RepositoryMySQL
}

func NewService(rmy RepositoryMySQL) Service {
	return &service{rmy}
}
