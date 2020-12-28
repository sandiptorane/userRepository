package test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepository/internal/database"
	"userRepository/internal/handlers"
	"userRepository/internal/token"
)

func TestProfile(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	userRepo := database.NewMockUserRepository(controller)
	handler := handlers.NewHandler(userRepo)

	r := mux.NewRouter()
	//first sign in
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`) //set login credential
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/signin", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	cookie := response.Result().Cookies()

	t.Run("Test for Get Profile:", func(t *testing.T) {
		r.HandleFunc("/profile", token.IsAuthorized(handler.GetProfile)).Methods("GET")
		req, err := http.NewRequest("GET", "/profile", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(cookie[0])
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
		assert.Equal(t, 200, response.Code, "ok response expected")
		assertGetProfile(t, response.Body.String())
	})
}

func assertGetProfile(t *testing.T,responseBody string){
	t.Helper()
	if responseBody == ""{
		t.Error("output is empty. It should be profile info or error info ")
	}
}

func TestUpdateProfile(t *testing.T){
	datastore,err := database.DbConnect()
	handler := handlers.NewHandler(datastore)

	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`)  //set login credential
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST","/signin",bytes.NewReader(body))
	if err!=nil{
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response,req)
	cookie := response.Result().Cookies()

	t.Run("test for update Profile",func(t *testing.T){
		r.HandleFunc("/profile", token.IsAuthorized(handler.UpdateProfile)).Methods("PUT")
		body := []byte( `{
		    "username" : "sandip123",
			"password" : "sandip@123",
			"lastname": "torane",
			"firstname" : "sandip",
			"age" : 23,
			"gender" :"male",
			"city" : "Ichalkaranji",
			"country" :"India",
			"phone" : "7972797852",
			"email":"sandip@gmail.com",
			"githubUsername" : "https://github.com/sandip"
		}`)

		// Create a request to pass to our handler.
		req, err := http.NewRequest("PUT", "/profile", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(cookie[0])
		response := httptest.NewRecorder()
		r.ServeHTTP(response,req)
		assert.Equal(t,200,response.Code,"ok response expected")
	})
}