package database

import (
	"errors"
	"fmt"
	"userRepository/internal/user"
)

type Register interface {
	AddUser(p *user.Person) error
	userAlreadyExists(username string) bool
}

var alreadyExistsError = "user is already present please choose another username"

//AddUser will register user's details in userRepository table and return error if occurs
func (repository *Datastore)AddUser(p *user.Person) error {
	if repository.userAlreadyExists(p.Username) {
		return errors.New(fmt.Sprintf("%s %s", p.Username, alreadyExistsError))
	}
	query := `INSERT INTO userRepository(username,password,firstname,lastname,age,gender,city,country,phone,email,githubUsername) VALUES
			(?,?,?,?,?,?,?,?,?,?,?)`
	//insert user's registration details into userRepository
	_, err := repository.Db.Exec(query, p.Username, p.Password, p.Firstname, p.Lastname, p.Age, p.Gender, p.City, p.Country, p.Phone, p.EmailId, p.GithubUsername)
	return err
}

//userAlreadyExists checks if user already present or not
func (repository *Datastore)userAlreadyExists(username string) bool {
	query := `SELECT username FROM userRepository WHERE username =?`
	var returnedUser string
	err := repository.Db.QueryRowx(query, username).Scan(&returnedUser)
	if err != nil && returnedUser == "" {
		return false
	}
	if username == returnedUser {
		return true
	}
	return false
}




