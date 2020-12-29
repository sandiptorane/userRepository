package hashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string)(hash string){
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes)
}
//func HashPassword(credential *user.Credential){

