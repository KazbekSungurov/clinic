package repository

import model "go_clinic/app/models"

type Repository struct {
	ScheduleDetails
	Appointment
	ScheduleOperation
	ScheduleConsult
	Department
}

type ScheduleDetails interface {
	Create(m *model.ScheduleDetails) (id uint, err error)
	Return(id uint) (m *model.ScheduleDetails, err error)
	Update(m *model.ScheduleDetails) (err error)
}

type Appointment interface {
	Create(m *model.Appointment) (id uint, err error)
	Return(id uint) (m *model.Appointment, err error)
	Update(m *model.Appointment) (err error)
	Delete(id uint) (err error)
}

type ScheduleOperation interface {
	Create(m *model.ScheduleOperation) (id uint, err error)
	Return(id uint) (m *model.ScheduleOperation, err error)
	Update(m *model.ScheduleOperation) (err error)
}

type ScheduleConsult interface {
	Create(m *model.ScheduleConsult) (id uint, err error)
	Return(id uint) (m *model.ScheduleConsult, err error)
	Update(m *model.ScheduleConsult) (err error)
}

type Department interface {
	Return(id uint) (m *model.Department, err error)
}
