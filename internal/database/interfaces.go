package database

//UserRepository defines all methods for database operations
type UserRepository interface{
	Register
	SignIn
	Profile
	GithubDetails
	Task
}
