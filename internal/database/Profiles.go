package database

import (
	log "github.com/sirupsen/logrus"
	"userRepository/internal/user"
)

//Profile interface wraps GetProfile and UpdateProfile methods
type Profile interface {
	GetProfile(username string)(*user.Person,error)
	UpdateProfile(p *user.Person)error
}

func newUser() user.Person {
	return user.Person{}
}

//GetProfile fetch user profile info from userRepository and return user.Person and error
func (repository *Datastore)GetProfile(username string)(*user.Person,error){
	person := newUser()  //initialize user.Person  and will used to store profile info
	query := `SELECT * FROM userRepository WHERE username = ?`
	err := repository.Db.Get(&person, query, username)  //get person profile details
	if err != nil {
		return nil, err
	}
	return &person, nil
}

//UpdateProfile updates the user profile and return error if any
func (repository *Datastore)UpdateProfile(p *user.Person)error {
	query := `UPDATE userRepository SET password=?,firstname=?,lastname=?,age=?,gender=?,city=?,country=?,phone=?,email=?,githubUsername=? WHERE username = ?`
	changes, err := repository.Db.Preparex(query)
	if err != nil {
		return err
	}
	_, err = changes.Exec(p.Password, p.Firstname, p.Lastname, p.Age, p.Gender, p.City, p.Country, p.Phone, p.EmailId, p.GithubUsername, p.Username)
	if err != nil {
		return err
	}
	log.Println("profile updated on database")
	return nil
}