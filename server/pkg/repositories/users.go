package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/uncleTen276/attendance.git/server/internal/postgresql"
	"github.com/uncleTen276/attendance.git/server/pkg/models"
)

type UserRepository interface {
	Create(u *models.CreateUser, ctx context.Context) error
	GetUserByIdentifier(identifier string) (*models.User, error)
	GetUserById(int) (*models.User, error)
	UpdatePasswordById(password string, id uuid.UUID) error
}

type UserRepo struct {
	db *postgresql.DB
}

const (
	UserTable = "users"
)

var _ UserRepository = (*UserRepo)(nil)

func NewUserRepo(db *postgresql.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) Create(u *models.CreateUser, ctx context.Context) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (identifier,password,first_name,last_name ) VALUES ($1,$2,$3,$4)`, UserTable)
	_, err := repo.db.Exec(query, u.Identifier, u.Password, u.FirstName, u.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) GetUserByIdentifier(identifier string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE identifier = $1`, UserTable)
	err := repo.db.Get(user, query, identifier)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) GetUserById(uId int) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, UserTable)
	err := repo.db.Get(user, query, uId)

	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, err
}

func (repo *UserRepo) UpdatePasswordById(password string, id uuid.UUID) error {
	query := fmt.Sprintf(`UPDATE "%s" SET password = $1 WHERE id = $2`, UserTable)
	_, err := repo.db.Exec(query, password, id)
	if err != nil {
		return err
	}

	return nil
}
