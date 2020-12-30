package database

import (
	log "github.com/sirupsen/logrus"
	"time"
	"userRepository/internal/tasks"
)

type Task interface {
	AddTask(task *tasks.Task,user string)
	GetTasks(username string,startDate time.Time,endDate time.Time)([]tasks.FilteredTasks,error)
	GetSingleTask(username string,id int,startDate time.Time,endDate time.Time)(*tasks.FilteredTasks,error)
}

func (repository *Datastore)AddTask(task *tasks.Task,user string){
	query := `INSERT INTO task(username,name,description,start,end,urlLink) VALUES(?,?,?,?,?,?)`
	_,err := repository.Db.Exec(query,user,task.Name,task.Description,task.Start,task.End,task.UrlLink)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	log.Println("task added successfully")
}


func (repository *Datastore)GetTasks(username string,startDate time.Time,endDate time.Time)([]tasks.FilteredTasks,error){
	var taskList []tasks.FilteredTasks
	formattedFrom := startDate.Format("2006-01-02 15:04:05")
	formattedTo := endDate.Format("2006-01-02 15:04:05")
	query := `SELECT id,name,description,start,end,urlLink FROM task WHERE (start BETWEEN ? AND ?) AND username = ?`
	err := repository.Db.Select(&taskList, query, formattedFrom, formattedTo, username)
	return taskList, err
}

func (repository *Datastore)GetSingleTask(username string,id int,startDate time.Time,endDate time.Time)(*tasks.FilteredTasks,error){
	var task tasks.FilteredTasks
	formattedFrom := startDate.Format("2006-01-02 15:04:05")
	formattedTo := endDate.Format("2006-01-02 15:04:05")
	query := `SELECT id,name,description,start,end,urlLink FROM task WHERE (start BETWEEN ? AND ?) AND (username = ? AND id = ?)`
	err := repository.Db.Get(&task, query, formattedFrom, formattedTo, username,id)
	return &task, err
}