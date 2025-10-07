package social

import "log"

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Test() error {
	log.Println("social repository Test() called")
	return nil
}
