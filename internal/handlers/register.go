package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/hashing"
	"userRepository/internal/user"
	"userRepository/pkg/validation"
)

//newUser initialize the user.Person
func newUser() *user.Person {
	return &user.Person{}
}

//Registration is endpoint to register new user to userRepository
func (handler *Handlers)Registration(w http.ResponseWriter,req *http.Request) {
	log.Println("new user registering to the system")

	w.Header().Set("Content-Type", "application/json")
	if req.Body == nil {
		fmt.Fprintln(w, "nil body passed")
		return
	}

	person := newUser()  //initialize the person
	err := json.NewDecoder(req.Body).Decode(&person)   //Decode person from json
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	validationError := validation.ValidateUser(person) //validate inputs of user and display errors if any
	if validationError != nil {
		validation.DisplayError(w, validationError)
		return
	}
    person.Password= hashing.HashPassword(person.Password)  //encrypt/hash password

	err = handler.Repository.AddUser(person) //store registration data into database userRepository
	if err != nil {
		fmt.Fprint(w, err.Error())
		log.Error(err.Error())
		return
	}

	fmt.Fprint(w, "new user registered successfully")
	log.Println("new user registered successfully")
}
