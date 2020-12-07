package database

import (
	log "github.com/sirupsen/logrus"
	"userRepository/internal/user"
)

func GetProfile(username string) (*user.Person,error){
	defer log.Println("returning profile details from database")
	  repository,_ := DbConnect()
      person := newUser()

      query := `SELECT * FROM userRepository WHERE username = ?`
      err := repository.Db.Get(&person,query,username)
      if err!=nil{
      	return nil,err
	  }
	  return &person,nil
}

func newUser() user.Person {
	return user.Person{}
}

func UpdateProfile(p *user.Person) error {
	repository, _ := DbConnect()
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
