package handlers

import (
	"net/http"
	"userRepository/internal/token"
)

//SignOut to signout current user
func SignOut(w http.ResponseWriter,req *http.Request) {
	token.ClearToken(w,req)
}