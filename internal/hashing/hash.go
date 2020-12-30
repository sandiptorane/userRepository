package hashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string)(hash string){
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes)
}

func VerifyPassword(hashedPassword,userPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(userPassword))
	if err==nil{
		return true
	}
	return false
}

