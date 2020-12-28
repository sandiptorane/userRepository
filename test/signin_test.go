package test

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userRepository/internal/database"
	"userRepository/internal/handlers"
)

type testCase struct {
	name string
	body []byte
	buildStubs func(userRepo *database.MockUserRepository)
	expectedStatusCode int
	expectedResponse string
}

func getSigninTestCases() []testCase{
	 testCases := []testCase{
		{
			name: "Test for signin successful",
			body: []byte(`{"username":"sandip123","password":"sandip@123"}`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().UserExists("sandip123", "sandip@123").Return(true)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "Welcome",
		},
		{
			name: "Test for unauthorised user",
			body: []byte(`{"username":"sandip123","password":"sandip@123"}`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().UserExists(gomock.Any(), gomock.Any()).Return(false)
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "You are Unauthorized to access the application.",
		},
	}
	return testCases
}

func signInRouter(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	return r
}

func TestSignInMock(t *testing.T){
	testCases := getSigninTestCases()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			userRepo := database.NewMockUserRepository(controller)
            tc.buildStubs(userRepo)
			req, err := http.NewRequest("POST", "/signin", bytes.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			response := httptest.NewRecorder()
			r := signInRouter(userRepo)  //mux router
			r.ServeHTTP(response, req)

			result := response.Result()
			if len(result.Cookies()) != 0 { //token stored in cookie
				cookie := result.Cookies()
				fmt.Println("cookie:", cookie[0])
			}

			checkStatus(t,tc.expectedStatusCode,response.Code)

			expectedResponse := tc.expectedResponse
			actualResponse := response.Body.String()
			assertBody(t, expectedResponse, actualResponse)
		})
	}
}

func checkStatus(t *testing.T,expectedCode int,actualCode int){
	t.Helper()
	if expectedCode!=actualCode{
		t.Error("expected status:",expectedCode,"but got :",actualCode)
	}
}

func assertBody(t *testing.T,expected string,actual string){
	t.Helper()
	if !strings.Contains(actual,expected){
		t.Errorf("unsuccessful:  output should contains %s\n but got %s\n",expected,actual)
	}
}

