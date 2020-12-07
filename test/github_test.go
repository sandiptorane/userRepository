package test

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userRepository/internal/handlers"
	"userRepository/internal/token"
)

func TestGetGithub(t *testing.T){
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
	result := response.Result()
    cookie := result.Cookies()

    //check get github
	r.HandleFunc("/github", token.IsAuthorized(handlers.Github)).Methods("GET")
	req, err = http.NewRequest("GET", "/github",nil)
	if err != nil {
		t.Fatal(err)
	}
	response = httptest.NewRecorder()
	req.AddCookie(cookie[0])
	r.ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"ok response expected")



	t.Run("if github account exists:",func(t *testing.T){
		expected := `githubUsername`
		actualOutput  := response.Body.String()
		assertGithub(t,expected,actualOutput)
	})

}

func assertGithub(t *testing.T,expected string,actualString string){
	t.Helper()
	if !strings.Contains(actualString,expected){
		t.Errorf("output should contains '%s' but got '%s'",expected,actualString)
	}
}

func TestGetGithubNotExists(t *testing.T){
	r := mux.NewRouter()

	//sign in first
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip1234","password":"sandip@123"}`)  //set login credential
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST","/signin",bytes.NewReader(body))
	if err!=nil{
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response,req)
	result := response.Result()
	cookie := result.Cookies()

	//check get github
	r.HandleFunc("/github", token.IsAuthorized(handlers.Github)).Methods("GET")
	req, err = http.NewRequest("GET", "/github",nil)
	if err != nil {
		t.Fatal(err)
	}
	response = httptest.NewRecorder()
	req.AddCookie(cookie[0])
	r.ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"ok response expected")

	t.Run("if github account doesn't exists:",func(t *testing.T){
		expected := `Github account doesn't exist please update the profile`
		actualOutput  := response.Body.String()
		assertGithub(t,expected,actualOutput)
	})
}