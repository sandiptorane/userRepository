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
	"userRepository/internal/tasks"
	"userRepository/internal/token"
)

func newGetTaskRouter(userRepo *database.MockUserRepository) *mux.Router{
	handler := handlers.NewHandler(userRepo)
	r := mux.NewRouter()
	r.HandleFunc("/signin", handler.SignIn).Methods("POST")
	r.HandleFunc("/task",token.IsAuthorized(handler.GetTasks)).Methods("GET")
	r.HandleFunc("/task/{id}",token.IsAuthorized(handler.GetSingleTask)).Methods("GET")
	return r
}

type getTaskTestCase struct {
	name string
	username string
	password string
	body []byte
	buildStubs func(userRepo *database.MockUserRepository)
	expectedStatusCode int
	expectedResponse string
}

func getTaskTestCases() [] getTaskTestCase{
	testCases := []getTaskTestCase{
		{
			name:     "Task displayed successfully",
			username: "sandip123",
			password: "sandip@123",
			body: []byte( `{
    		      "startDate": "2020-12-03",
    		      "endDate": "2020-12-05"
		          }`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				username := "sandip123"
				returnTasks := []tasks.FilteredTasks{
					{
						Id: 1,
						Task: tasks.Task{
							Name:        "project meeting",
							Description: "project discussion",
							Start:       "2020-12-04 2:15:03",
							End:         "2020-12-04 3:00:00",
							UrlLink:     "https://meet.google.com/sdjljire",
						},
					},
					{
						Id: 4,
						Task: tasks.Task{
							Name:        "project meeting2",
							Description: "project discussion",
							Start:       "2020-12-05 1:25:00",
							End:         "2020-12-05 2:30:00",
							UrlLink:     "https://meet.google.com/dkjirm",
						},
					},
				}
				userRepo.EXPECT().GetTasks(username, gomock.Any(), gomock.Any()).Return(returnTasks, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintln(`[{"Id":1,"name":"project meeting","description":"project discussion","start":"2020-12-04 2:15:03","end":"2020-12-04 3:00:00","urlLink":"https://meet.google.com/sdjljire"},{"Id":4,"name":"project meeting2","description":"project discussion","start":"2020-12-05 1:25:00","end":"2020-12-05 2:30:00","urlLink":"https://meet.google.com/dkjirm"}]`),
			},
			{
			name:     "No tasks assigned",
			username: "sandip123",
			password: "sandip@123",
			body: []byte( `{
    		"startDate": "2020-12-03",
    		"endDate": "2020-12-03"
			}`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().GetTasks("sandip123",gomock.Any(),gomock.Any()).Return(nil,nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("No tasks assigned\n"),
		},
	}
	return testCases
}

func TestGetTasks(t *testing.T){
	testCases := getTaskTestCases()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			userRepo := database.NewMockUserRepository(controller)
			router := newGetTaskRouter(userRepo)  //mux router
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

			//AddTasks handler
			tc.buildStubs(userRepo)
			req, err = http.NewRequest("GET", "/task",bytes.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",auth)
			response = httptest.NewRecorder()
			router.ServeHTTP(response,req)

			checkStatus(t,tc.expectedStatusCode,response.Code)

			expectedResponse := tc.expectedResponse
			actualResponse := response.Body.String()
			checkGetTasksResponse(t, expectedResponse, actualResponse)
		})
	}
}

func checkGetTasksResponse(t *testing.T,expected string,actualString string){
	t.Helper()
	if actualString!=expected{
		t.Errorf("\n output excpected '%s' \n but got '%s'",expected,actualString)
	}
}

func TestGetSingleTask(t *testing.T){     //when pass tasks id in url then display particular task
	testCases := getSingleTaskTestCases()
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			userRepo := database.NewMockUserRepository(controller)
			router := newGetTaskRouter(userRepo)  //mux router
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

			//AddTasks handler
			tc.buildStubs(userRepo)
			req, err = http.NewRequest("GET", "/task/1",bytes.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",auth)
			response = httptest.NewRecorder()
			router.ServeHTTP(response,req)

			checkStatus(t,tc.expectedStatusCode,response.Code)

			expectedResponse := tc.expectedResponse
			actualResponse := response.Body.String()
			checkGetTasksResponse(t, expectedResponse, actualResponse)
		})
	}
}

func getSingleTaskTestCases() [] getTaskTestCase{
	testCases := []getTaskTestCase{
		{
			name:     "Task displayed successfully",
			username: "sandip123",
			password: "sandip@123",
			body: []byte( `{
    		      "startDate": "2020-12-03",
    		      "endDate": "2020-12-05"
		          }`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				username := "sandip123"
				returnTasks := &tasks.FilteredTasks{
						Id: 1,
						Task: tasks.Task{
							Name:        "project meeting",
							Description: "project discussion",
							Start:       "2020-12-04 2:15:03",
							End:         "2020-12-04 3:00:00",
							UrlLink:     "https://meet.google.com/sdjljire",
						},
					}
				userRepo.EXPECT().GetSingleTask(username,gomock.Any(), gomock.Any(), gomock.Any()).Return(returnTasks, nil)
				},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintln(`{"Id":1,"name":"project meeting","description":"project discussion","start":"2020-12-04 2:15:03","end":"2020-12-04 3:00:00","urlLink":"https://meet.google.com/sdjljire"}`),
		},
		{
			name:     "No tasks assigned",
			username: "sandip123",
			password: "sandip@123",
			body: []byte( `{
    		"startDate": "2020-12-03",
    		"endDate": "2020-12-03"
			}`),
			buildStubs: func(userRepo *database.MockUserRepository) {
				userRepo.EXPECT().GetSingleTask("sandip123",gomock.Any(),gomock.Any(),gomock.Any()).Return(&tasks.FilteredTasks{},errors.New("no rows found"))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   fmt.Sprintf("No tasks assigned\n"),
		},
	}
	return testCases
}






