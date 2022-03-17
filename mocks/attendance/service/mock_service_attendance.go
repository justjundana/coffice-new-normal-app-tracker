// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_attendanceService is a generated GoMock package.
package mock_attendanceService

import (
	reflect "reflect"
	attendanceEntities "sirclo/project-capstone/entities/attendanceEntities"
	attendanceRequest "sirclo/project-capstone/utils/request/attendanceRequest"

	gomock "github.com/golang/mock/gomock"
)

// MockAttServiceInterface is a mock of AttServiceInterface interface.
type MockAttServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAttServiceInterfaceMockRecorder
}

// MockAttServiceInterfaceMockRecorder is the mock recorder for MockAttServiceInterface.
type MockAttServiceInterfaceMockRecorder struct {
	mock *MockAttServiceInterface
}

// NewMockAttServiceInterface creates a new mock instance.
func NewMockAttServiceInterface(ctrl *gomock.Controller) *MockAttServiceInterface {
	mock := &MockAttServiceInterface{ctrl: ctrl}
	mock.recorder = &MockAttServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttServiceInterface) EXPECT() *MockAttServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateAttendance mocks base method.
func (m *MockAttServiceInterface) CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAttendance", loginId, input)
	ret0, _ := ret[0].(attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAttendance indicates an expected call of CreateAttendance.
func (mr *MockAttServiceInterfaceMockRecorder) CreateAttendance(loginId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAttendance", reflect.TypeOf((*MockAttServiceInterface)(nil).CreateAttendance), loginId, input)
}

// GetAttendancesById mocks base method.
func (m *MockAttServiceInterface) GetAttendancesById(attID string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendancesById", attID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAttendancesById indicates an expected call of GetAttendancesById.
func (mr *MockAttServiceInterfaceMockRecorder) GetAttendancesById(attID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendancesById", reflect.TypeOf((*MockAttServiceInterface)(nil).GetAttendancesById), attID)
}

// GetAttendancesCurrentUser mocks base method.
func (m *MockAttServiceInterface) GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendancesCurrentUser", userId, status, order)
	ret0, _ := ret[0].([]attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendancesCurrentUser indicates an expected call of GetAttendancesCurrentUser.
func (mr *MockAttServiceInterfaceMockRecorder) GetAttendancesCurrentUser(userId, status, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendancesCurrentUser", reflect.TypeOf((*MockAttServiceInterface)(nil).GetAttendancesCurrentUser), userId, status, order)
}

// GetAttendancesRangeDate mocks base method.
func (m *MockAttServiceInterface) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order string) ([]attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttendancesRangeDate", employeeEmail, dateStart, dateEnd, status, officeId, order)
	ret0, _ := ret[0].([]attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttendancesRangeDate indicates an expected call of GetAttendancesRangeDate.
func (mr *MockAttServiceInterfaceMockRecorder) GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, officeId, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttendancesRangeDate", reflect.TypeOf((*MockAttServiceInterface)(nil).GetAttendancesRangeDate), employeeEmail, dateStart, dateEnd, status, officeId, order)
}

// IsCheckins mocks base method.
func (m *MockAttServiceInterface) IsCheckins() ([]attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCheckins")
	ret0, _ := ret[0].([]attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCheckins indicates an expected call of IsCheckins.
func (mr *MockAttServiceInterfaceMockRecorder) IsCheckins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCheckins", reflect.TypeOf((*MockAttServiceInterface)(nil).IsCheckins))
}

// UpdateAttendance mocks base method.
func (m *MockAttServiceInterface) UpdateAttendance(loginId string, input attendanceRequest.UpdateAttRequest) (attendanceEntities.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttendance", loginId, input)
	ret0, _ := ret[0].(attendanceEntities.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAttendance indicates an expected call of UpdateAttendance.
func (mr *MockAttServiceInterfaceMockRecorder) UpdateAttendance(loginId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttendance", reflect.TypeOf((*MockAttServiceInterface)(nil).UpdateAttendance), loginId, input)
}
