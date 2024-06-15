package models

import "time"

type Events struct {
	Id         int       `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"-"`
	UpdatedAt  time.Time `db:"updated_at" json:"-"`
	Hash       string    `db:"hash" json:"hash_tx"`
	Type       string    `db:"type_tx" json:"type_tx"`
	EmployeeId int       `db:"employee_id" json:"employee_id"`
}

type CreateEvent struct {
	Hash       string `db:"hash" json:"hash_tx"`
	Type       string `db:"type_tx" json:"type_tx"`
	EmployeeId int    `db:"employee_id" json:"employee_id"`
}
