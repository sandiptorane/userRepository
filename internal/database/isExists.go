package database

import "log"

func UserExists(username,password string) bool{
	repository,err := DbConnect()
	if err !=nil{
		return false
	}
	query := `SELECT username FROM userRepository WHERE username =? AND password = ?`
	var returnedUser string
	err = repository.Db.QueryRowx(query,username,password).Scan(&returnedUser)
	if err != nil {
		log.Println("in UserExists of isExists err:",err)
		return false
	}
	if username == returnedUser{
		return true
	}
	return false
}


