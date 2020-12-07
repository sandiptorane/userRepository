package test

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepository/internal/handlers"
	"userRepository/internal/token"
)

func TestAddTask(t *testing.T) {
	r := mux.NewRouter()

	//sign in first
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`) //set login credential
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/signin", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
   cookie := response.Result().Cookies()

	t.Run("Test for addTask:", func(t *testing.T) {
		r.HandleFunc("/task", handlers.AddTasks).Methods("POST")
		body = []byte(`{
    	"name": "project",
    	"description": "discussion on project",
		"start":"2020-11-22 10:25:20",
    	"end" : "2020-11-22 11:25:10",`)
		req, err := http.NewRequest("POST", "/task", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(cookie[0])
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
	})
}

func TestGetTask(t *testing.T) {
	r := mux.NewRouter()

	//sign in first
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`)  //set login credential
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST","/signin",bytes.NewReader(body))
	if err!=nil{
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response,req)
	cookie := response.Result().Cookies()

	t.Run("Test for getTask:",func(t *testing.T){
		r.HandleFunc("/task", token.IsAuthorized(handlers.GetTasks)).Methods("GET")
		body = []byte(`{"startdate":"2020-11-21",
						"endDate" : "2020-11-22"
			}`)
		req, err := http.NewRequest("GET","/task",bytes.NewReader(body))
		if err!=nil{
			t.Fatal(err)
		}
		req.AddCookie(cookie[0])
		response := httptest.NewRecorder()
		r.ServeHTTP(response,req)
		assert.Equal(t,http.StatusOK,response.Code,"ok expected")
	})
}