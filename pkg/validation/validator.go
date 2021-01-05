package validation

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
	"userRepository/internal/tasks"
	"userRepository/internal/user"
)

//ValidateUser validates inputs for validate rule
//validate rules defined in user.Person struct
func ValidateUser(person *user.Person) error{
	log.Println("validate details entered by user")
	v := validator.New()
	err := v.Struct(person)
	return err
}
//ValidateCredential validates inputs for validate rule
//validate rules defined in user.Credential struct while user signing
func ValidateCredential(credential *user.Credential)error{
	v := validator.New()
	err := v.Struct(credential)
	return err
}

//ValidateTask will validate task inputs
//validate rules defined in tasks.Task struct
func ValidateTask(task *tasks.Task) error{
	log.Println("validate task")
	v := validator.New()
	err := v.Struct(task)
	return err
}

//ValidateTime checks startTask and endTask. it should be format of "YYYY-MM-DD Hr:min:sec"
//this same format is used in database storage. if format not correct return error
func ValidateTime(startTask string, endTask string) error{
	const (
		layoutUS = "2006-01-02 15:04:05"    //layout is as per standard linux Mon Jan 2 15:04:05 -0700 MST 2006
		dateTimeError = "start and end task format should be : YYYY-MM-DD Hr:min:sec"
	)

	_,startTaskErr := time.Parse(layoutUS,startTask)
	_,endTaskErr := time.Parse(layoutUS,endTask)
	if startTaskErr != nil || endTaskErr !=nil{
		return errors.New(dateTimeError)
	}
	return nil
}
//ValidateDate checks startDate and endDate. it should be format of "YYYY-MM-DD"
//it returns dates in time.Time format and error if any
func ValidateDate(startDate, endDate string)(time.Time,time.Time,error){
	start,err := validateDate(startDate) //
	if err!=nil{
		return start,time.Time{},err
	}

	end,err := validateDate(endDate)
	if err!=nil{
		return start,end,err
	}
	return start,end,nil
}

//validateDate will parse the date for "YYYY-MM-DD" layout and return it in time.Time format and error if any
func validateDate(taskDate string) (time.Time, error){
	log.Println("validating date for format YYYY-MM-DD")
	const (
		dateLayout = "2006-01-02"
		dateError = "StartDate and endDate format should be : YYYY-MM-DD"
	)

	date,err := time.Parse(dateLayout,taskDate)
	if err!=nil{
		return date,errors.New(dateError)
	}
	return date,nil
}

//display validation errors
func DisplayError(w http.ResponseWriter,errs error){
	for _, e := range errs.(validator.ValidationErrors){
		fmt.Fprintln(w,e)
	}
}