package database

import (
	"errors"
	"fmt"
	"userRepository/internal/user"
)

var alreadyExistsError = "user is already present please choose another username"

//InsertUserinfo will register user's details in userRepository table
func InsertUserinfo(p *user.Person) (err error){
	repository,err := DbConnect()
	if err!=nil{
		return err
	}
	//repository.CreateUserRepository() //It will create userRepository table if not exists in the database
	err = repository.addUser(p)
	return err
}

func (repository *Datastore)addUser(p *user.Person) error{
	if repository.userAlreadyExists(p.Username){
		return errors.New(fmt.Sprintf("%s %s",p.Username, alreadyExistsError))
	}
	query := `INSERT INTO userRepository(username,password,firstname,lastname,age,gender,city,country,phone,email,githubUsername) VALUES
			(?,?,?,?,?,?,?,?,?,?,?)`
	_,err := repository.Db.Exec(query,p.Username,p.Password,p.Firstname,p.Lastname,p.Age,p.Gender,p.City,p.Country,p.Phone,p.EmailId,p.GithubUsername)
   return err
}

func (dbInstance *Datastore)userAlreadyExists(username string) bool{
	query := `SELECT username FROM userRepository WHERE username =?`
	var returnedUser string
	err := dbInstance.Db.QueryRowx(query,username).Scan(&returnedUser)
	if err != nil && returnedUser == ""{
		return false
	}
	if username == returnedUser{
		return true
	}
	return false
}




