package tasks

//DateFiltering to store start Date and end date to get task
type DateFiltering struct {
	StartDate string `json:"startdate" validate:"required"` //format= YYYY-MM-DD
	EndDate   string `json:"enddate" validate:"required"`   //format= YYYY-MM-DD
}

