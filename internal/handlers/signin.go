package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/database"
	"userRepository/internal/token"
)

type Credential struct{
	Username string `json:"username"`
	Password string `json:"password"`
}


//SignIn endpoint will allow user to enter in the system
func SignIn(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type","application/json")
	var creds Credential
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !userExists(creds.Username,creds.Password){
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are Unauthorized to access the application.\n"))
		return
	}
	token.CreateToken(creds.Username,w,req)
	fmt.Fprintf(w,"Welcome %s !\n",creds.Username)
	log.Println("user successfully signed in")
}

func userExists(username string,password string)bool{
	log.Println("in userExists of signin")
	if database.UserExists(username,password){
		return true
	}
	return false
}
