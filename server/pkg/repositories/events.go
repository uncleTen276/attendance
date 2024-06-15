package repositories

import (
	"fmt"

	"github.com/uncleTen276/attendance.git/server/internal/postgresql"
	"github.com/uncleTen276/attendance.git/server/pkg/models"
)

type EventsRepository interface {
	Create(*models.CreateEvent) error
}

type EventRepo struct {
	db *postgresql.DB
}

const (
	EventTable   = "events"
	CREATE_EVENT = "create"
	UPDATE_EVENT = "update"
)

var _ EventsRepository = (*EventRepo)(nil)

func NewEventRepo(db *postgresql.DB) EventsRepository {
	return &EventRepo{db}
}

func (repo *EventRepo) Create(e *models.CreateEvent) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (hash, type_tx, employee_id) VALUES ($1,$2,$3)`, EventTable)
	_, err := repo.db.Exec(query, e.Hash, e.Type, e.EmployeeId)
	if err != nil {
		return err
	}
	return nil
}
