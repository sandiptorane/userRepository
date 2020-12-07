package database

import (
	log "github.com/sirupsen/logrus"
	"time"
	"userRepository/internal/tasks"
)

func AddTask(task *tasks.Task,user string){
	repository,_ := DbConnect()
	//repository.CreateTasks()   // it will create task table in database if not exists
	query := `INSERT INTO task(username,name,description,start,end,urlLink) VALUES(?,?,?,?,?,?)`
	_,err := repository.Db.Exec(query,user,task.Name,task.Description,task.Start,task.End,task.UrlLink)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	log.Println("task added successfully")
}

func GetTasks(username string,startDate time.Time,endDate time.Time)([]tasks.Task,error) {
	repository, _ := DbConnect()
	var taskList []tasks.Task
	formattedFrom := startDate.Format("2006-01-02 15:04:05")
	formattedTo := endDate.Format("2006-01-02 15:04:05")
	query := `SELECT name,description,start,end FROM task WHERE (start BETWEEN ? AND ?) AND username = ?`
	err := repository.Db.Select(&taskList, query, formattedFrom, formattedTo, username)
	return taskList, err
}