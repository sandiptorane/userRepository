package handlers

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
    "userRepository/internal/tasks"
    "userRepository/pkg/token"
	"userRepository/pkg/validation"
)

//initialize the tasks.DateFiltering to get StartDate and endDate
func newFilter() *tasks.DateFiltering {
      return &tasks.DateFiltering{}
}

//display current user's all tasks from startDate to endDate
func (handler *Handlers)GetTasks(w http.ResponseWriter,req *http.Request) {
    log.Println("getting task list of current user from startDate to endDate")
    w.Header().Set("Content-Type", "application/json")
    username := token.GetUserName(w, req)

    filter := newFilter()    //initialize dateFilter
    err := json.NewDecoder(req.Body).Decode(&filter)   //parse startDate and endDate from response body
    if err != nil {
        fmt.Fprintln(w, err.Error())
        return
    }

    startDate, endDate, dateError := validation.ValidateDate(filter.StartDate, filter.EndDate)  //validate Date
    if dateError != nil {
        fmt.Fprintln(w, dateError)
        return
    }

    //get all tasks from database of current user from startDate to endDate
    taskList, err := handler.Repository.GetTasks(username, startDate, endDate)
    if err != nil {
        fmt.Fprintln(w, err)
        return
    }
    if taskList == nil {
        fmt.Fprintln(w, "No tasks assigned")
        return
    }
    err = json.NewEncoder(w).Encode(taskList) //display task
    if err == nil {
        log.Println("task displayed")
    }
}

//display current user's particular single task from startDate to endDate
//URL input: /task/{id}
func (handler *Handlers)GetSingleTask(w http.ResponseWriter,req *http.Request) {
    log.Println("getting task list of current user from startDate to endDate")
    w.Header().Set("Content-Type", "application/json")
    username := token.GetUserName(w, req)
    filter := newFilter()
    err := json.NewDecoder(req.Body).Decode(&filter)
    if err != nil {
        fmt.Fprintln(w, "json parsing error:", err.Error())
        return
    }
    startDate, endDate, dateError := validation.ValidateDate(filter.StartDate, filter.EndDate)
    if dateError != nil {
        fmt.Fprintln(w, dateError)
        return
    }

    params := mux.Vars(req)  //routes variable
    id, _ := strconv.Atoi(params["id"])   //fetch id from params

    //get task from database for particular id
    singleTask, err := handler.Repository.GetSingleTask(username, id, startDate, endDate)
    if err != nil {
        fmt.Fprintln(w, "No tasks assigned")
        return
    }
    err = json.NewEncoder(w).Encode(singleTask) //display task
    if err == nil {
        log.Println("task displayed")
    }
}