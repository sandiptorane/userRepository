package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/user"
	"userRepository/pkg/token"
	"userRepository/pkg/validation"
)

//profile contains person info
type profile struct{
   person user.Person
}

//updateProfile will update the profile of current existent user
func (handler *Handlers)UpdateProfile(w http.ResponseWriter,req *http.Request) {
	log.Println("updating current user's profile")
	w.Header().Set("Content-Type", "application/json")
	if req.Body == nil {
		fmt.Fprintln(w, "nil body passed")
		return
	}
	p := profile{
		person: user.Person{},
	}
	err := json.NewDecoder(req.Body).Decode(&p.person)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	validationError := validation.ValidateUser(&p.person) //validate inputs of person
	if validationError != nil {
		validation.DisplayError(w, validationError)
		return
	}

	username := token.GetUserName(w, req)  //get username from token
	if !checkUsername(username, p.person.Username) { //check current user name and entered username are same. if not same display error
		fmt.Fprintln(w, "you can,t update username please enter existing username")
		return
	}

	//update profile on the database
	err = handler.Repository.UpdateProfile(&p.person)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "Profile updated successfully")
	log.Println("Profile updated successfully")
}

//verify actual username and entered username because user can't update its username
func checkUsername(actualUsername string,enteredUsername string)bool {
	if actualUsername != enteredUsername {
		return false
	}
	return true
}