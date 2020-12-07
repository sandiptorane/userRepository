package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/database"
	"userRepository/internal/user"
	"userRepository/internal/validation"
)

func newUser() *user.Person {
	return &user.Person{}
}

//Registration is endpoint to register new user to userRepository
func Registration(w http.ResponseWriter,req *http.Request){
	log.Println("new user registering to the system")

	w.Header().Set("Content-Type","application/json")
	if req.Body == nil{
		fmt.Fprintln(w,"nil body passed")
		return
	}

    person := newUser()
	err := json.NewDecoder(req.Body).Decode(&person)
	if err!=nil{
		fmt.Fprintln(w,err.Error())
		return
	}

	validationError := validation.ValidateUser(person) //validate inputs of user
	if validationError!=nil{
		validation.DisplayError(w,validationError)
		return
	}

	err = database.InsertUserinfo(person) //store registration data into database userRepository
	if err!=nil{
		fmt.Fprintln(w,err.Error())
		log.Error(err.Error())
		return
	}

	fmt.Fprintln(w,"new user registered successfully")
	log.Println("new user registered successfully")
}
