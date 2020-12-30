package test

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userRepository/internal/database"
	"userRepository/internal/handlers"
	"userRepository/internal/token"
)

func getGithubRouter(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/github",token.IsAuthorized(handler.Github)).Methods("GET")
	return r
}

type githubTestCase struct {
	name string
	username string
	password string
	buildStubs func(userRepo *database.MockUserRepository)
	expectedStatusCode int
	expectedResponse string
}

func getGithubTestCases() []githubTestCase{
	testCases := []githubTestCase{
		{
			name: "If github account exists",
			username: "sandip123",
			password: "sandip@123",
			buildStubs: func(userRepo *database.MockUserRepository) {
				username := "sandip123"
				githubUsername := "https://github.com/sandip123"
				userRepo.EXPECT().GetGithub(username).Return(githubUsername)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"githubUsername":"https://github.com/sandip123"}`,
		},
		{
			name : "If github account doesn't exists",
			username : "sandip123",
			password: "sandip@123",
			buildStubs: func(userRepo *database.MockUserRepository) {
				username := "sandip123"
				githubUsername := ""
				userRepo.EXPECT().GetGithub(username).Return(githubUsername)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `Github account doesn't exist please update the profile`,
		},
	}
	return testCases
}

func TestGithub(t *testing.T){
	testCases := getGithubTestCases()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			userRepo := database.NewMockUserRepository(controller)

			//first sign in
			userRepo.EXPECT().UserExists(tc.username, tc.password).Return(true)
			body := []byte(`{"username":"`+tc.username+`","password":"`+tc.password+`"}`)
			req, err := http.NewRequest("POST", "/signin", bytes.NewReader(body))
			if err != nil {
				t.Fatal(err)
			}
			response := httptest.NewRecorder()
			router := getGithubRouter(userRepo)  //mux router
			router.ServeHTTP(response, req)

			auth := req.Header.Get("Authorization")  //jwt auth token stored in the Bearer token Authorization header

			//Github handler
			tc.buildStubs(userRepo)
			req, err = http.NewRequest("GET", "/github",nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",auth)
			response = httptest.NewRecorder()
			router.ServeHTTP(response,req)

			checkStatus(t,tc.expectedStatusCode,response.Code)

			expectedResponse := tc.expectedResponse
			actualResponse := response.Body.String()
			assertGithub(t, expectedResponse, actualResponse)
		})
	}
}

func assertGithub(t *testing.T,expected string,actualString string){
	t.Helper()
	if !strings.Contains(actualString,expected){
		t.Errorf("output should contains '%s' but got '%s'",expected,actualString)
	}
}


