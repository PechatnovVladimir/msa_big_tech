package chat

import "log"

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) Test() error {
	log.Println("chat repository Test() called")
	return nil
}
