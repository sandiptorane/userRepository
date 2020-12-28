package database

type UserRepository interface{
	Register
	SignIn
	Profile
	GithubDetails
	Task
}
