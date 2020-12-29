package handlers

import "userRepository/internal/database"

type Handlers struct {
	Repository database.UserRepository
}

func NewHandler(repository database.UserRepository) *Handlers{
	return &Handlers{
		Repository: repository,
	}
}
