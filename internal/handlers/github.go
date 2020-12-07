package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/database"
	"userRepository/internal/token"
)

type gitHub interface {
	getGithub(username string)
}

type githubInfo struct {
	GithubUsername string `json:"githubUsername"`
}

func newGithub() *githubInfo {
	return &githubInfo{}
}

//Github will print github accounts details
func Github(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type","application/json")
	userName := token.GetUserName(w,req)
	log.Println(userName)
	log.Println("getting github details")
	github := newGithub()
	github.getGithub(userName)
	if github.GithubUsername == ""{
		fmt.Fprintln(w,"Github account doesn't exist please update the profile")
		return
	}
	err := json.NewEncoder(w).Encode(github)
	if err != nil{
		fmt.Fprintln(w,err)
	}
}

func (g *githubInfo)getGithub(username string){
	g.GithubUsername = database.GetGithub(username)
}
