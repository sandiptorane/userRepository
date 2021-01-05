package hashing

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)
//HashPassword to encrypt password during user's Registration
func HashPassword(password string)(hash string){
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes)
}

//VerifyPassword compare hashed password of database and actual password and return true if same
func VerifyPassword(hashedPassword,userPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(userPassword))
	if err==nil{
		return true
	}
	log.Print("hash compare err:",err)
	return false
}

