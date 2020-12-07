package tasks

type DateFiltering struct{
	StartDate string  `json:"startdate" validate:"required"`     //format= YYYY-MM-DD
	EndDate   string   `json:"enddate" validate:"required"`      //format= YYYY-MM-DD
}
