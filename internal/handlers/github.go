package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/pkg/token"
)

//githubInfo contains github account details
type githubInfo struct {
	GithubUsername string `json:"githubUsername"`
}

func newGithub() *githubInfo {
	return &githubInfo{}
}

//Github will print github accounts details og logged user
func (handler *Handlers)Github(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("getting github details")
	username := token.GetUserName(w, req)
	github := newGithub()

	//get github account details from database
	github.GithubUsername = handler.Repository.GetGithub(username)
	if github.GithubUsername == "" {
		fmt.Fprintln(w, "Github account doesn't exist please update the profile")
		return
	}

	err := json.NewEncoder(w).Encode(github)   //display github account details
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

