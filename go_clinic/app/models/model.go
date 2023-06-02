package model

import "time"

type PersonalData struct {
}

type User struct {
}

type Patient struct {
}

type Employee struct {
}

type ScheduleDetails struct {
	ID                  uint           `json:"id,omitempty"`
	WeekDays            []time.Weekday `json:"week_days,omitempty"`
	StartAt             time.Time      `json:"start_at"`
	EndAt               time.Time      `json:"end_at"`
	ExaminationDuration time.Duration  `json:"examination_duration,omitempty"`
	BreakStart          time.Time      `json:"break_start"`
	BreakDuration       time.Duration  `json:"break_duration,omitempty"`
	AppointmentsLimit   int            `json:"appointments_limit,omitempty"`
}

type Appointment struct {
	ID       uint      `json:"id,omitempty"`
	Day      time.Time `json:"day"`
	Patients []Patient `json:"patients,omitempty"`
}

type ScheduleOperation struct {
	ID              uint          `json:"id,omitempty"`
	Appointments    []Appointment `json:"appointments,omitempty"`
	ScheduleDetails `json:"schedule_details"`
}

type ScheduleConsult struct {
	ID              uint          `json:"id,omitempty"`
	Appointments    []Appointment `json:"appointments,omitempty"`
	ScheduleDetails `json:"schedule_details"`
}

type Department struct {
	ID                uint       `json:"id,omitempty"`
	Name              string     `json:"name,omitempty"`
	Employees         []Employee `json:"employees,omitempty"`
	ScheduleConsult   `json:"schedule_consult"`
	ScheduleOperation `json:"schedule_operation"`
}
