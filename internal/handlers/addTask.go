package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/tasks"
	"userRepository/pkg/token"
	"userRepository/pkg/validation"
)

//add current user's task
func (handler *Handlers)AddTasks(w http.ResponseWriter,req *http.Request) {
	username := token.GetUserName(w, req)
	task := new(tasks.Task)
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	validationError := validation.ValidateTask(task)  //validate task inputs
	if validationError != nil {
		validation.DisplayError(w, validationError)
		return
	}
	if err := validation.ValidateTime(task.Start, task.End); err != nil {   //validate time for YYYY-MM-DD hh:mm:sec format
		fmt.Fprintln(w, err)
		return
	}

	//AddTask to database
	handler.Repository.AddTask(task, username)
	fmt.Fprintln(w, "Task added")
	log.Println("Task added by the user")
}
