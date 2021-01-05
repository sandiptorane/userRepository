package database

import (
	"log"
)

type GithubDetails interface {
	GetGithub(username string) string
}

//GetGithub return user's github Account details
func (repository *Datastore)GetGithub(username string) string {
	query := `SELECT githubUsername FROM userRepository WHERE username = ?`
	var github string
	err := repository.Db.Get(&github, query, username)
	if err != nil {
		log.Println("error:", err)
	}
	return github
}
