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

type tokenString struct {
	Token string `json:"token"`
}

//SignIn endpoint will allow user to enter in the system
func (handler *Handlers)SignIn(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds *user.Credential
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if validationError := validation.ValidateCredential(creds); validationError != nil { //validate inputs of user
		validation.DisplayError(w, validationError)
		return
	}
	if !handler.Repository.UserExists(creds.Username, creds.Password) { //check user exists or not in database
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are Unauthorized to access the application.\n"))
		log.Println("sign in unsuccessful")
		return
	}

	//Create token for signed user
	tokenStr, err := token.CreateToken(creds.Username, w,req)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	//display created token
	json.NewEncoder(w).Encode(tokenString{Token: tokenStr,})

	log.Println("user successfully signed in")
}