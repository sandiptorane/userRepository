package handlers

import "userRepository/internal/database"

//Handlers holds datastore connection and api endpoint Methods
type Handlers struct {
	Repository database.UserRepository
}

//NewHandler initialize the Handlers
func NewHandler(repository database.UserRepository) *Handlers{
	return &Handlers{
		Repository: repository,
	}
}
