package tasks

//Task contains following details to add task
type Task struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Start       string `json:"start" validate:"required"` //format= YYYY-MM-DD Hr:min:sec
	End         string `json:"end" validate:"required"`   //format= YYYY-MM-DD Hr:min:sec
	UrlLink     string `json:"urlLink,omitempty" validate:"omitempty,url" db:"urlLink"`
}


//FilteredTasks to store task details returned from the database
type FilteredTasks struct {
	Id int  //Id is auto incremented in the database
	Task
}
