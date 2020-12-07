package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/database"
	"userRepository/internal/handlers"
	"userRepository/internal/token"
	"userRepository/internal/vipers"
)

//test connection with database. if not able to connect terminate the application
func init(){
	_,err := database.DbConnect()
	if err!=nil{
		log.Fatal("not able connect with server database: application terminated")
	}
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Registration).Methods("POST")
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	r.HandleFunc("/profile", token.IsAuthorized(handlers.GetProfile)).Methods("GET")
	r.HandleFunc("/profile", token.IsAuthorized(handlers.UpdateProfile)).Methods("PUT")
	r.HandleFunc("/github", token.IsAuthorized(handlers.Github)).Methods("GET")
	r.HandleFunc("/task", token.IsAuthorized(handlers.AddTasks)).Methods("POST")
	r.HandleFunc("/task", token.IsAuthorized(handlers.GetTasks)).Methods("GET")
	r.HandleFunc("/signout", token.IsAuthorized(handlers.SignOut)).Methods("POST")
	port := vipers.GetPort()
	log.Fatal(http.ListenAndServe(port,r))
}
