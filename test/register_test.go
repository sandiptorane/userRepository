package test

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userRepository/internal/database"
	"userRepository/internal/handlers"
)

func router(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/register", handler.Registration).Methods("POST")
	return r
}

func TestRegister(t *testing.T){
	var body []byte
	body = []byte( `{
		"username" : "sandip123",
			"password" : "sandip@123",
			"lastname": "torane",
			"firstname" : "sandip",
			"aGe" : 22,
			"gender" :"male",
			"city" : "Ichalkaranji",
			"country" :"India",
			"phone" : "7972797852",
			"email":"sandip@gmail.com"
		}`)
	//write test cases
	testCases := []struct {
		name string
		body []byte
		buildStubs func(userRepo *database.MockUserRepository)
		expectedResponse string

	}{
		{
			name: "Registration successful",
			body: body,
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().AddUser(gomock.Any()).Return(nil)
			},
			expectedResponse: "new user registered successfully",
		},
		{
			name : "test for unsuccessful registration",
			body: body,
			buildStubs: func(userRepo *database.MockUserRepository){
				userRepo.EXPECT().AddUser(gomock.Any()).Return(errors.New("user is already present please choose another username"))
			},
			expectedResponse: "user is already present please choose another username",
		},
	}

	for i:= range testCases{
	tc :=testCases[i]
	 t.Run(tc.name,func(t *testing.T) {
		 controller := gomock.NewController(t)
		 userRepo := database.NewMockUserRepository(controller)
		 tc.buildStubs(userRepo)
		 // Create a request to pass to our handler.
		 req, err := http.NewRequest("POST", "/register", bytes.NewReader(tc.body))
		 if err != nil {
			 t.Fatal(err)
		 }
		 response := httptest.NewRecorder()
		 router(userRepo).ServeHTTP(response, req)
		 assert.Equal(t, 200, response.Code, "ok response expected")
		 expectedResponse := tc.expectedResponse
		 actualResponse := response.Body.String()
		 checkResponse(t, actualResponse, expectedResponse)
	 })
	}
}

func checkResponse(t *testing.T,actualResponse string,expectedResponse string){
	t.Helper()
	if !strings.Contains(actualResponse,expectedResponse){
		t.Error("expected:",expectedResponse,"but actual response :",actualResponse)
	}
}
