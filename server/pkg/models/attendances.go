package models

type AttendanceCheckIn struct {
	Position string `json:"position" validate:"required"`
	CheckIn  string `json:"check_in" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Detail   string `json:"detail"`
}

type AttendaceUpdate struct {
	Position string `json:"position" validate:"required"`
	CheckIn  string `json:"check_in" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	RecordId int    `json:"record_id" validate:"required,number"`
	Detail   string `json:"detail"`
}
