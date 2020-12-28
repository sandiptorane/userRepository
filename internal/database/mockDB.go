// Code generated by MockGen. DO NOT EDIT.
// Source: userRepository/internal/database (interfaces: UserRepository)

// Package database is a generated GoMock package.
package database

import (
	reflect "reflect"
	time "time"
	tasks "userRepository/internal/tasks"
	user "userRepository/internal/user"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockUserRepository) AddTask(arg0 *tasks.Task, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTask", arg0, arg1)
}

// AddTask indicates an expected call of AddTask.
func (mr *MockUserRepositoryMockRecorder) AddTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockUserRepository)(nil).AddTask), arg0, arg1)
}

// AddUser mocks base method.
func (m *MockUserRepository) AddUser(arg0 *user.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserRepositoryMockRecorder) AddUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserRepository)(nil).AddUser), arg0)
}

// GetGithub mocks base method.
func (m *MockUserRepository) GetGithub(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGithub", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGithub indicates an expected call of GetGithub.
func (mr *MockUserRepositoryMockRecorder) GetGithub(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGithub", reflect.TypeOf((*MockUserRepository)(nil).GetGithub), arg0)
}

// GetProfile mocks base method.
func (m *MockUserRepository) GetProfile(arg0 string) (*user.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0)
	ret0, _ := ret[0].(*user.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockUserRepositoryMockRecorder) GetProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockUserRepository)(nil).GetProfile), arg0)
}

// GetSingleTask mocks base method.
func (m *MockUserRepository) GetSingleTask(arg0 string, arg1 int, arg2, arg3 time.Time) (*tasks.FilteredTasks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSingleTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*tasks.FilteredTasks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSingleTask indicates an expected call of GetSingleTask.
func (mr *MockUserRepositoryMockRecorder) GetSingleTask(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSingleTask", reflect.TypeOf((*MockUserRepository)(nil).GetSingleTask), arg0, arg1, arg2, arg3)
}

// GetTasks mocks base method.
func (m *MockUserRepository) GetTasks(arg0 string, arg1, arg2 time.Time) ([]tasks.FilteredTasks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]tasks.FilteredTasks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockUserRepositoryMockRecorder) GetTasks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockUserRepository)(nil).GetTasks), arg0, arg1, arg2)
}

// UpdateProfile mocks base method.
func (m *MockUserRepository) UpdateProfile(arg0 *user.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserRepositoryMockRecorder) UpdateProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserRepository)(nil).UpdateProfile), arg0)
}

// UserExists mocks base method.
func (m *MockUserRepository) UserExists(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UserExists indicates an expected call of UserExists.
func (mr *MockUserRepositoryMockRecorder) UserExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserExists", reflect.TypeOf((*MockUserRepository)(nil).UserExists), arg0, arg1)
}

// userAlreadyExists mocks base method.
func (m *MockUserRepository) userAlreadyExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "userAlreadyExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// userAlreadyExists indicates an expected call of userAlreadyExists.
func (mr *MockUserRepositoryMockRecorder) userAlreadyExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "userAlreadyExists", reflect.TypeOf((*MockUserRepository)(nil).userAlreadyExists), arg0)
}
