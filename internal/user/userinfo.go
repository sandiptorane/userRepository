package user

//Person contains following details for registration
type Person struct {
	Username string  	`json:"username" validate:"required"`
	Password string  	`json:"password" validate:"required"`
	Firstname string  	`json:"firstname" validate:"required"`
	Lastname  string  	`json:"lastname" validate:"required"`
	Age int           	`json:"age" validate:"min=0,max=130,required"`
	Gender string     	`json:"gender" validate:"oneof=male female"`
	City string       	`json:"city" validate:"required"`
	Country string 	  	`json:"country,omitempty"`
	Phone string 	    `json:"phone" validate:"len=10,numeric"`
	EmailId string 		`json:"email" validate:"required,email" db:"email"`
	GithubUsername string `json:"githubUsername,omitempty" validate:"omitempty,startswith=https://github.com/" db:"githubUsername"`
}



