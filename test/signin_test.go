package test

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userRepository/internal/handlers"
)

func TestSignIn(t *testing.T){
	r := mux.NewRouter()
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`)
	req, err := http.NewRequest("POST","/signin",bytes.NewReader(body))
	if err!=nil{
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response,req)

	result:= response.Result()
	cookie := result.Cookies()
	fmt.Println("cookie:",cookie[0])

	t.Run("check SignIn endpoint's status:",func(t *testing.T){
		assert.Equal(t,200,response.Code,"ok response expected")
	})

	t.Run("user validation:",func(t *testing.T) {
		expected := "Welcome"
		actual := response.Body.String()
		assertBody(t, expected, actual)
	})
}

func assertBody(t *testing.T,expected string,actual string){
	if !strings.Contains(actual,expected){
		t.Errorf("unsuccessful:  output should contains %s\n but got %s\n",expected,actual)
	}
}
