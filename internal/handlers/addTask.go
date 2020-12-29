package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"userRepository/internal/tasks"
	"userRepository/internal/token"
	"userRepository/internal/validation"
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
	validationError := validation.ValidateTask(task)
	if validationError != nil {
		validation.DisplayError(w, validationError)
		return
	}
	if err := validation.ValidateTime(task.Start, task.End); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	handler.Repository.AddTask(task, username)
	fmt.Fprintln(w, "Task added")
	log.Println("Task added by the user")
}
