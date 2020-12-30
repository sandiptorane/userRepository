package test

import (
	"bytes"
	"errors"
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

func getUpdateProfileRouter(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/profile",token.IsAuthorized(handler.UpdateProfile)).Methods("PUT")
	return r
}

type updateProfileTestCase struct {
	name string
	username string
	password string
	body []byte
	buildStubs func(userRepo *database.MockUserRepository)
	expectedStatusCode int
	expectedResponse string
}

func getUpdateProfileTestCases() [] updateProfileTestCase{
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
	testCases := []updateProfileTestCase{
		{
			name:     "Successful update",
			username: "sandip123",
			password: "sandip@123",
			body: body,
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().UpdateProfile(gomock.Any()).Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("Profile updated successfully\n"),
		},
		{
			name:     "don't update username",
			username: "sandip123",
			password: "sandip@123",
			//enter wrong username
			body: []byte( `{
		    "username" : "pankaj123",
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
			}`),
			buildStubs: func(userRepo *database.MockUserRepository) {

			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("you can,t update username please enter existing username\n"),
		},
		{
			name:     "test for update error",
			username: "sandip123",
			password: "sandip@123",
			body: body,
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().UpdateProfile(gomock.Any()).Return(errors.New("something went wrong not able to update"))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("something went wrong not able to update\n"),
		},
	}
	return testCases
}

func TestUpdateProfileData(t *testing.T){
	testCases := getUpdateProfileTestCases()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			userRepo := database.NewMockUserRepository(controller)
			router := getUpdateProfileRouter(userRepo)  //mux router
			//first sign in
			userRepo.EXPECT().UserExists(tc.username, tc.password).Return(true)
			body := []byte(`{"username":"`+tc.username+`","password":"`+tc.password+`"}`)
			req, err := http.NewRequest("POST", "/signin", bytes.NewReader(body))
			if err != nil {
				t.Fatal(err)
			}
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)

			auth := req.Header.Get("Authorization")  //jwt auth token stored in the Bearer token Authorization header

			//GetProfile handler
			tc.buildStubs(userRepo)
			req, err = http.NewRequest("PUT", "/profile",bytes.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",auth)
			response = httptest.NewRecorder()
			router.ServeHTTP(response,req)

			checkStatus(t,tc.expectedStatusCode,response.Code)

			expectedResponse := tc.expectedResponse
			actualResponse := response.Body.String()
			checkUpdateProfileResponse(t, expectedResponse, actualResponse)
		})
	}
}

func checkUpdateProfileResponse(t *testing.T,expected string,actualString string){
	t.Helper()
	if actualString!=expected{
		t.Errorf("\n output excpected '%s' \n but got '%s'",expected,actualString)
	}
}




