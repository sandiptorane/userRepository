package test

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepository/internal/handlers"
)

func router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/register", handlers.Registration).Methods("POST")
	return r
}

func TestRegister(t *testing.T){
	var body []byte
	body = []byte( `{
		"username" : "sandip123",
			"password" : "sandip@123",
			"lastname": "torane",
			"firstname" : "sandip",
			"aGe" : 23,
			"gender" :"male",
			"city" : "Ichalkaranji",
			"country" :"India",
			"phone" : "7972797852",
			"email":"sandip@gmail.com"

	}`)
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	router().ServeHTTP(response,req)
	assert.Equal(t,200,response.Code,"ok response expected")
}
