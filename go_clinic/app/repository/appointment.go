package repository

import (
	"database/sql"
	model "go_clinic/app/models"
)

type appointment struct {
	db *sql.DB
}

func (a *appointment) Create(m *model.Appointment) (id uint, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *appointment) Return(id uint) (m *model.Appointment, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *appointment) Update(m *model.Appointment) (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *appointment) Delete(id uint) (err error) {
	//TODO implement me
	panic("implement me")
}
