package database

import (
	"log"
)
type CreateOperations interface {
	CreateUserRepository()
	CreateTasks()
}

//CreateUserRepository will create table to store user's details
func (dbInstance *Datastore)CreateUserRepository(){
	query := `CREATE TABLE IF NOT EXISTS userRepository(
				username VARCHAR(50),
				password TEXT NOT NULL,
                firstname TEXT NOT NULL,
                lastname TEXT NOT NULL,
                age INTEGER NOT NULL,
   				gender TEXT NOT NULL,
   				city TEXT NOT NULL,
				country TEXT NULL,
				phone VARCHAR(10) NOT NULL,
				email TEXT NOT NULL,
				githubUsername TEXT NULL,
				PRIMARY KEY (username)
				);`

	_,err := dbInstance.Db.Exec(query)
	if err!=nil{
		log.Fatalln(err)
	}
	//log.Println("userRepository table created successfully : if not exists")
}

//CreateTasks will create table to store user's task
func (dbInstance *Datastore)CreateTasks(){
	query := `CREATE TABLE IF NOT EXISTS task(
		username VARCHAR(50) NOT NULL,
		name TEXT NOT NULL,
    	description TEXT,
    	start datetime NOT NULL,
    	end   datetime NOT NULL,
    	urlLink   TEXT NULL 
		);`
	_,err := dbInstance.Db.Exec(query)
	if err!=nil{
		log.Fatalln(err)
	}
	//log.Println("tasks table created successfully")
}

