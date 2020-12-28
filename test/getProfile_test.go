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
	"userRepository/internal/token"
	"userRepository/internal/user"
)

func getUpdateProfileRouter(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/profile",token.IsAuthorized(handler.GetProfile)).Methods("GET")
	return r
}

type getProfileTestCase struct {
	name string
	username string
	password string
	buildStubs func(userRepo *database.MockUserRepository)
	expectedStatusCode int
	expectedResponse string
}

func getProfileTestCases() []getProfileTestCase{
	testCases := []getProfileTestCase{
		{
			name: "Successful",
			username: "sandip123",
			password: "sandip@123",
			buildStubs: func(userRepo *database.MockUserRepository) {
				username := "sandip123"
				userRepo.EXPECT().GetProfile(username).Return(&user.Person{
									Username:       "sandip123",
									Password:       "sandip@123",
									Firstname:      "sandip",
									Lastname:       "torane",
									Age:            22,
									Gender:         "male",
									City:           "Ichalkaranji",
									Country:        "India",
									Phone:          "7945867158",
									EmailId:        "sandip@gmail.com",
									GithubUsername: "https://github.com/sandip",},nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse: `{"username":"sandip123","password":"sandip@123","firstname":"sandip","lastname":"torane","age":22,"gender":"male","city":"Ichalkaranji","country":"India","phone":"7945867158","email":"sandip@gmail.com","githubUsername":"https://github.com/sandip"}`,
		},

	}
	return testCases
}

func TestGetProfile(t *testing.T){
	testCases := getProfileTestCases()
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
			router := getUpdateProfileRouter(userRepo)  //mux router
			router.ServeHTTP(response, req)

			result := response.Result()
			cookie := result.Cookies()

			//GetProfile handler
			tc.buildStubs(userRepo)
			req, err = http.NewRequest("GET", "/profile",nil)
			if err != nil {
				t.Fatal(err)
			}
			response = httptest.NewRecorder()
			if len(cookie)!=0 {
				req.AddCookie(cookie[0])
			}
			router.ServeHTTP(response,req)
			checkStatus(t,tc.expectedStatusCode,response.Code)

			expectedResponse := tc.expectedResponse
			actualResponse := response.Body.String()
			checkGetProfileResponse(t, expectedResponse, actualResponse)
		})
	}
}

func checkGetProfileResponse(t *testing.T,expected string,actualString string){
	t.Helper()
	fmt.Printf("expected: %T",expected)
	fmt.Printf("actual : %T",actualString)
	if !strings.Contains(actualString,expected){
		t.Errorf("\n output excpected '%s' \n but got '%s':",expected,actualString)

	}
}



