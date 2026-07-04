package auth

type Repository interface {
	
}

type Validator struct {
	repository Repository
}

func New (repository Repository) Validator {
	return Validator{
		repository: repository,
	}
}