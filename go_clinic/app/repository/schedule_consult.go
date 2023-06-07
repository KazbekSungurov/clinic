package repository

import (
	model "go_clinic/app/models"
)

type scheduleConsult struct{}

func (s *scheduleConsult) Create(m *model.ScheduleConsult) (id uint, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *scheduleConsult) Return(id uint) (m *model.ScheduleConsult, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *scheduleConsult) Update(m *model.ScheduleConsult) (err error) {
	//TODO implement me
	panic("implement me")
}
