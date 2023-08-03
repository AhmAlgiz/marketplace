package repository

type Auth interface {
}

type Item interface {
}

type Repository struct {
	Auth
	Item
}

func NewRepository() *Repository {
	return &Repository{}
}
