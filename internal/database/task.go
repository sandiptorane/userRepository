package database

import (
	log "github.com/sirupsen/logrus"
	"time"
	"userRepository/internal/tasks"
)
//Task wraps methods: AddTask, GetTasks, GetSingleTask
type Task interface {
	AddTask(task *tasks.Task,user string)
	GetTasks(username string,startDate time.Time,endDate time.Time)([]tasks.FilteredTasks,error)
	GetSingleTask(username string,id int,startDate time.Time,endDate time.Time)(*tasks.FilteredTasks,error)
}

//AddTask add current user's task to database
func (repository *Datastore)AddTask(task *tasks.Task,user string){
	query := `INSERT INTO task(username,name,description,start,end,urlLink) VALUES(?,?,?,?,?,?)`
	_,err := repository.Db.Exec(query,user,task.Name,task.Description,task.Start,task.End,task.UrlLink)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	log.Println("task added successfully")
}

//GetTasks fetch all filtered tasks of user between startDate to endDate and return all tasks  and error
func (repository *Datastore)GetTasks(username string,startDate time.Time,endDate time.Time)([]tasks.FilteredTasks,error){
	var taskList []tasks.FilteredTasks    //to store all returned task in TaskList
	formattedFrom := startDate.Format("2006-01-02 15:04:05")  //format startDate to "YYYY-MM-DD hrs:min:sec"
	formattedTo := endDate.Format("2006-01-02 15:04:05")     	//format endDate to "YYYY-MM-DD hrs:min:sec"
	query := `SELECT id,name,description,start,end,urlLink FROM task WHERE (start BETWEEN ? AND ?) AND username = ?`
	err := repository.Db.Select(&taskList, query, formattedFrom, formattedTo, username)
	return taskList, err
}

//GetSingleTask fetch particular filtered task of user using {id} between startDate to endDate  and return task and error
func (repository *Datastore)GetSingleTask(username string,id int,startDate time.Time,endDate time.Time)(*tasks.FilteredTasks,error){
	var task tasks.FilteredTasks      //task to store returned task
	formattedFrom := startDate.Format("2006-01-02 15:04:05")  //format startDate to "YYYY-MM-DD hrs:min:sec"
	formattedTo := endDate.Format("2006-01-02 15:04:05")		//format endDate to "YYYY-MM-DD hrs:min:sec"
	query := `SELECT id,name,description,start,end,urlLink FROM task WHERE (start BETWEEN ? AND ?) AND (username = ? AND id = ?)`
	err := repository.Db.Get(&task, query, formattedFrom, formattedTo, username,id)
	return &task, err
}