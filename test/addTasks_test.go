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

func getAddTaskRouter(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/task",token.IsAuthorized(handler.AddTasks)).Methods("POST")
	return r
}

type addTaskTestCase struct {
	name string
	username string
	password string
	body []byte
	buildStubs func(userRepo *database.MockUserRepository)
	expectedStatusCode int
	expectedResponse string
}

func getAddTaskTestCases() [] addTaskTestCase{
	testCases := []addTaskTestCase{
		{
			name:     "Task Added successfully",
			username: "sandip123",
			password: "sandip@123",
			body: []byte( `{
			"name" :"project meeting",
    		"description" : "discussion",
    		"start": "2020-12-03 02:54:24",
    		"end": "2020-12-03 03:57:24",
			"urlLink" : "https://meet.google.com/sdjljire"
		}`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().AddTask(gomock.Any(),"sandip123")
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("Task added\n"),
		},
		{
			name:     "Validate start and end time format",
			username: "sandip123",
			password: "sandip@123",
			body: []byte( `{
			"name" :"project meeting",
    		"description" : "discussion",
    		"start": "2020-12-03 02:24",
    		"end": "2020-12-03 03:57:24"
		}`),
			buildStubs: func(userRepo *database.MockUserRepository) {
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("start and end task format should be : YYYY-MM-DD Hr:min:sec\n"),
		},
	}
	return testCases
}

func TestAddTask(t *testing.T){
	testCases := getAddTaskTestCases()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			userRepo := database.NewMockUserRepository(controller)
			router := getAddTaskRouter(userRepo)  //mux router
			//first sign in
			userRepo.EXPECT().UserExists(tc.username, tc.password).Return(true)
			body := []byte(`{"username":"`+tc.username+`","password":"`+tc.password+`"}`)
			req, err := http.NewRequest("POST", "/signin", bytes.NewReader(body))
			if err != nil {
				t.Fatal(err)
			}
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)

			result := response.Result()
			cookie := result.Cookies()  //jwt auth token stored in the cookie

			//AddTasks handler
			tc.buildStubs(userRepo)
			req, err = http.NewRequest("POST", "/task",bytes.NewReader(tc.body))
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
			checkAddTaskResponse(t, expectedResponse, actualResponse)
		})
	}
}

func checkAddTaskResponse(t *testing.T,expected string,actualString string){
	t.Helper()
	if actualString!=expected{
		t.Errorf("\n output excpected '%s' \n but got '%s'",expected,actualString)
	}
}





