package database

import (
	"log"
)

type GithubDetails interface {
	getGithub(username string) string
}

func GetGithub(username string) string{
	repository,_ := DbConnect()
	defer repository.Db.Close()
	return repository.getGithub(username)
}

func (repository *Datastore)getGithub(username string) string{
	query := `SELECT githubUsername FROM userRepository WHERE username = ?`
	var github string
	err := repository.Db.Get(&github,query,username)
	if err!=nil {
		log.Println("error:", err)
	}
	return github
}
