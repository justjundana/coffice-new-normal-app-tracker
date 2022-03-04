package attendanceService

import (
	"sirclo/project-capstone/entities/attendanceEntities"
	"sirclo/project-capstone/utils/request/attendanceRequest"
)

type AttServiceInterface interface {
	GetAttendances(employee, date, status, office, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesRangeDate(employeeEmail, dateStart, dateEnd, status, office, order string) ([]attendanceEntities.Attendance, error)
	GetAttendancesCurrentUser(userId, status, order string) ([]attendanceEntities.Attendance, error)
	CreateAttendance(loginId string, input attendanceRequest.CreateAttRequest) (attendanceEntities.Attendance, error)
	UpdateAttendance(loginId string, input attendanceRequest.UpdateAttRequest) (attendanceEntities.Attendance, error)
	CheckUserRole(loginId string) string
}
