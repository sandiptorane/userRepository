package database

import (
	"log"
	"userRepository/internal/hashing"
)

type SignIn interface {
	UserExists(username ,password string) bool
}


//UserExists check user is present or not in the database for signed user
func (repository *Datastore)UserExists(username,password string) bool {
	query := `SELECT password FROM userRepository WHERE username =?`
	var returnedPassword string
	//get password for user if user exists
	err := repository.Db.QueryRowx(query, username).Scan(&returnedPassword)
	if err != nil {
		log.Println("in UserExists of isExists err:", err)
		return false
	}
	//returnedPassword is hashed password. so verify the hashed password and actual password and return result
	if hashing.VerifyPassword(returnedPassword,password) {
		log.Println("user present in database")
		return true
	}
	log.Println("user is not present in database")
	return false
}