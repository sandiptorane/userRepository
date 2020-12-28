package database

import (
	log "github.com/sirupsen/logrus"
	"userRepository/internal/user"
)

type Profile interface {
	GetProfile(username string)(*user.Person,error)
	UpdateProfile(p *user.Person)error
}

func newUser() user.Person {
	return user.Person{}
}

func (repository *Datastore)GetProfile(username string)(*user.Person,error){
	person := newUser()
	query := `SELECT * FROM userRepository WHERE username = ?`
	err := repository.Db.Get(&person, query, username)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (repository *Datastore)UpdateProfile(p *user.Person)error{
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