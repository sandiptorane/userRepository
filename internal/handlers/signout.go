package handlers

import (
	"net/http"
	"userRepository/pkg/token"
)

//SignOut to signout current user
func (handler *Handlers)SignOut(w http.ResponseWriter,req *http.Request) {
     w.Header().Set("Content-Type","application/json")
	token.ClearToken(w, req)  //delete generated token
}