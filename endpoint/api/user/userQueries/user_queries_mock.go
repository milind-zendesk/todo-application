// Code generated by MockGen. DO NOT EDIT.
// Source: todo-application/endpoint/api/user/userQueries (interfaces: UserQueries)

// Package userqueries is a generated GoMock package.
package userqueries

import (
	reflect "reflect"
	model "todo-application/model"

	gomock "github.com/golang/mock/gomock"
)

// MockUserQueries is a mock of UserQueries interface.
type MockUserQueries struct {
	ctrl     *gomock.Controller
	recorder *MockUserQueriesMockRecorder
}

// MockUserQueriesMockRecorder is the mock recorder for MockUserQueries.
type MockUserQueriesMockRecorder struct {
	mock *MockUserQueries
}

// NewMockUserQueries creates a new mock instance.
func NewMockUserQueries(ctrl *gomock.Controller) *MockUserQueries {
	mock := &MockUserQueries{ctrl: ctrl}
	mock.recorder = &MockUserQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserQueries) EXPECT() *MockUserQueriesMockRecorder {
	return m.recorder
}

// GetAllUsersData mocks base method.
func (m *MockUserQueries) GetAllUsersData() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsersData")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsersData indicates an expected call of GetAllUsersData.
func (mr *MockUserQueriesMockRecorder) GetAllUsersData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsersData", reflect.TypeOf((*MockUserQueries)(nil).GetAllUsersData))
}

// GetUserData mocks base method.
func (m *MockUserQueries) GetUserData(arg0 int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserData", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserData indicates an expected call of GetUserData.
func (mr *MockUserQueriesMockRecorder) GetUserData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserData", reflect.TypeOf((*MockUserQueries)(nil).GetUserData), arg0)
}

// GetUserTodosData mocks base method.
func (m *MockUserQueries) GetUserTodosData(arg0 int) (model.UserTodoDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTodosData", arg0)
	ret0, _ := ret[0].(model.UserTodoDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTodosData indicates an expected call of GetUserTodosData.
func (mr *MockUserQueriesMockRecorder) GetUserTodosData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTodosData", reflect.TypeOf((*MockUserQueries)(nil).GetUserTodosData), arg0)
}

// InsertUserData mocks base method.
func (m *MockUserQueries) InsertUserData(arg0 model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUserData indicates an expected call of InsertUserData.
func (mr *MockUserQueriesMockRecorder) InsertUserData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserData", reflect.TypeOf((*MockUserQueries)(nil).InsertUserData), arg0)
}

// UpdateUserData mocks base method.
func (m *MockUserQueries) UpdateUserData(arg0 int, arg1 model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserData indicates an expected call of UpdateUserData.
func (mr *MockUserQueriesMockRecorder) UpdateUserData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserData", reflect.TypeOf((*MockUserQueries)(nil).UpdateUserData), arg0, arg1)
}
