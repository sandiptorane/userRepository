package handlers

import (
    "encoding/json"
    "fmt"
    log "github.com/sirupsen/logrus"
    "net/http"
    "userRepository/internal/database"
    "userRepository/internal/tasks"
    "userRepository/internal/token"
    "userRepository/internal/validation"
)

type userTask interface {
    addTask(w http.ResponseWriter,req *http.Request,username string)
    getTasks(w http.ResponseWriter,req *http.Request,username string)
}

type task struct {
    userTask
    tasks.Task
}

//add current user's task
func AddTasks(w http.ResponseWriter,req *http.Request){
      username := token.GetUserName(w,req)
      task := new(tasks.Task)
      err := json.NewDecoder(req.Body).Decode(&task)
      if err != nil {
            fmt.Fprintln(w, err.Error())
            return
      }
      validationError := validation.ValidateTask(task)
      if validationError != nil {
            validation.DisplayError(w, validationError)
            return
      }
      if err := validation.ValidateTime(task.Start, task.End); err != nil {
            fmt.Fprintln(w, err)
            return
      }
      database.AddTask(task,username)
      fmt.Fprintln(w,"Task added")
}


func newFilter() *tasks.DateFiltering {
      return &tasks.DateFiltering{}
}

//display current user's all tasks from startDate to endDate
func GetTasks(w http.ResponseWriter,req *http.Request){
      log.Println("getting task list of current user from startDate to endDate")
      w.Header().Set("Content-Type","application/json")
      username := token.GetUserName(w,req)
     filter := newFilter()
     err := json.NewDecoder(req.Body).Decode(&filter)
     if err!=nil{
           fmt.Fprintln(w,err.Error())
           return
     }
     startDate,endDate,dateError:= validation.ValidateDate(filter.StartDate,filter.EndDate)
     if dateError!=nil{
           fmt.Fprintln(w,dateError)
            return
     }
     taskList,err := database.GetTasks(username,startDate,endDate)
     if err!=nil{
           fmt.Fprintln(w,err)
           return
     }
     if taskList == nil{
        fmt.Fprintln(w,"No tasks assigned")
         return
     }
      err = json.NewEncoder(w).Encode(taskList)   //display task
      if err==nil{
          log.Println("task displayed")
      }
}