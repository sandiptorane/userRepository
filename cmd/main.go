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

func main() {
	datastore, err := database.DbConnect()
	if err != nil {        //if not able to connect with database then terminate the application
		log.Fatal("not able connect with server database: application terminated")
	}
	handler := handlers.NewHandler(datastore)

	r := mux.NewRouter()
	r.HandleFunc("/register", handler.Registration).Methods("POST")
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/profile", token.IsAuthorized(handler.GetProfile)).Methods("GET")
	r.HandleFunc("/profile", token.IsAuthorized(handler.UpdateProfile)).Methods("PUT")
	r.HandleFunc("/github", token.IsAuthorized(handler.Github)).Methods("GET")
	r.HandleFunc("/task", token.IsAuthorized(handler.AddTasks)).Methods("POST")
	r.HandleFunc("/task", token.IsAuthorized(handler.GetTasks)).Methods("GET")
	r.HandleFunc("/task/{id}", token.IsAuthorized(handler.GetSingleTask)).Methods("GET")
	r.HandleFunc("/signout", token.IsAuthorized(handler.SignOut)).Methods("POST")
	port := vipers.GetPort()
	log.Fatal(http.ListenAndServe(port, r))

	//close database
	datastore.Db.Close()
}