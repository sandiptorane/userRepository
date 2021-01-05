package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/pkg/token"
)

//get profile info and display
func (handler *Handlers)GetProfile(w http.ResponseWriter,req *http.Request) {
	log.Println("getting profile details")
	w.Header().Set("Content-Type", "application/json")
	username := token.GetUserName(w, req)
	//Get profile info from database
	person, err := handler.Repository.GetProfile(username)
	if err != nil {
		log.Println(err)
	}
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
}
