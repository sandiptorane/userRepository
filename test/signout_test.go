package test

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepository/internal/handlers"
)

func TestSignOut(t *testing.T){
	r := mux.NewRouter()
	r.HandleFunc("/signin", handlers.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`)
	req, err := http.NewRequest("POST","/signin",bytes.NewReader(body))
	if err!=nil{
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response,req)
	cookie := response.Result().Cookies()

	t.Run("Test for sign out:",func(t *testing.T){
		r.HandleFunc("/signout", handlers.SignOut).Methods("POST")
		req, err := http.NewRequest("POST","/signout",nil)
		if err!=nil{
			t.Fatal(err)
		}
		req.AddCookie(cookie[0])
		response := httptest.NewRecorder()
		r.ServeHTTP(response,req)
		fmt.Println(response.Body.String())
	})
}
