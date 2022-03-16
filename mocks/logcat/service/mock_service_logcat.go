// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_logcatService is a generated GoMock package.
package mock_logcatService

import (
	reflect "reflect"
	logcatEntities "sirclo/project-capstone/entities/logcatEntities"

	gomock "github.com/golang/mock/gomock"
)

// MockLogcatServiceInterface is a mock of LogcatServiceInterface interface.
type MockLogcatServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLogcatServiceInterfaceMockRecorder
}

// MockLogcatServiceInterfaceMockRecorder is the mock recorder for MockLogcatServiceInterface.
type MockLogcatServiceInterfaceMockRecorder struct {
	mock *MockLogcatServiceInterface
}

// NewMockLogcatServiceInterface creates a new mock instance.
func NewMockLogcatServiceInterface(ctrl *gomock.Controller) *MockLogcatServiceInterface {
	mock := &MockLogcatServiceInterface{ctrl: ctrl}
	mock.recorder = &MockLogcatServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogcatServiceInterface) EXPECT() *MockLogcatServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateLogcat mocks base method.
func (m *MockLogcatServiceInterface) CreateLogcat(loginId, message, category string) (logcatEntities.Logcat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogcat", loginId, message, category)
	ret0, _ := ret[0].(logcatEntities.Logcat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLogcat indicates an expected call of CreateLogcat.
func (mr *MockLogcatServiceInterfaceMockRecorder) CreateLogcat(loginId, message, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogcat", reflect.TypeOf((*MockLogcatServiceInterface)(nil).CreateLogcat), loginId, message, category)
}

// GetLogcatUser mocks base method.
func (m *MockLogcatServiceInterface) GetLogcatUser(userID string) ([]logcatEntities.Logcat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogcatUser", userID)
	ret0, _ := ret[0].([]logcatEntities.Logcat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogcatUser indicates an expected call of GetLogcatUser.
func (mr *MockLogcatServiceInterfaceMockRecorder) GetLogcatUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogcatUser", reflect.TypeOf((*MockLogcatServiceInterface)(nil).GetLogcatUser), userID)
}

// GetLogcats mocks base method.
func (m *MockLogcatServiceInterface) GetLogcats() ([]logcatEntities.Logcat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogcats")
	ret0, _ := ret[0].([]logcatEntities.Logcat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogcats indicates an expected call of GetLogcats.
func (mr *MockLogcatServiceInterfaceMockRecorder) GetLogcats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogcats", reflect.TypeOf((*MockLogcatServiceInterface)(nil).GetLogcats))
}