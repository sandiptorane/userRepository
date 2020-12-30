package test

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"userRepository/internal/database"
	"userRepository/internal/handlers"
	"userRepository/internal/token"
)

func TestSignOut(t *testing.T){
	controller := gomock.NewController(t)
	userRepo := database.NewMockUserRepository(controller)
	userRepo.EXPECT().UserExists("sandip123","sandip@123").Return(true)
	handler := handlers.NewHandler(userRepo)

	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	body := []byte(`{"username":"sandip123","password":"sandip@123"}`)
	req, err := http.NewRequest("POST","/signin",bytes.NewReader(body))
	if err!=nil{
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response,req)

	auth := req.Header.Get("Authorization")  //jwt auth token stored in the Bearer token Authorization header

	t.Run("Test for sign out:",func(t *testing.T){
		r.HandleFunc("/signout", token.IsAuthorized(handler.SignOut)).Methods("POST")
		req, err := http.NewRequest("POST","/signout",nil)
		if err!=nil{
			t.Fatal(err)
		}
		req.Header.Set("Authorization",auth)
		response := httptest.NewRecorder()
		r.ServeHTTP(response,req)

		actual := response.Body.String()
		expected := fmt.Sprintf("signed out successfully\n")
		checkStatus(t,http.StatusOK,response.Code)
		checkSignoutResponse(t,expected,actual)
	})
}

func checkSignoutResponse(t *testing.T,expected string,actual string){
	t.Helper()
	if expected!=actual{
		t.Error("wont:",expected,"but got:",actual)
	}
}
