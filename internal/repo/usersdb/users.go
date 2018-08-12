package usersdb

// usersdb is an lower level interface to database

import (
	"context"
	"database/sql"
	"time"

	"github.com/ges-sh/dbug/dbugdb"
)

// User is an exact representation of database users row
type User struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Pass      []byte    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
}

// CreateUser created new user with provided email and password.
func CreateUser(ctx context.Context, db dbugdb.DB, login string, password []byte) (User, error) {
	var u User
	err := db.QueryRowContext(ctx, `
	INSERT INTO users(login, pass)
		VALUES($1, $2)
		RETURNING id, login, pass, created_at, updated_at, status, role
	`, login, password).Scan(
		&u.ID,
		&u.Login,
		&u.Pass,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Status,
		&u.Role,
	)
	return u, err
}

func scanUser(row *sql.Row) (User, error) {
	var u User
	err := row.Scan(
		&u.ID,
		&u.Login,
		&u.Pass,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
		&u.Status,
		&u.Role,
	)
	return u, err
}

// FetchUserByID returns user based on provided id
func FetchUserByID(ctx context.Context, db dbugdb.DB, id int64) (User, error) {
	row := db.QueryRowContext(ctx, `
		SELECT
				id,
				login,
				pass,
				created_at,
				updated_at,
				deleted_at,
				status,
				role
			FROM users
			WHERE login = $1
		`)
	return scanUser(row)
}

// FetchUserByLogin returns user based on provided login
func FetchUserByLogin(ctx context.Context, db dbugdb.DB, login string) (User, error) {
	row := db.QueryRowContext(ctx, `
		SELECT
				id,
				login,
				pass,
				created_at,
				updated_at,
				deleted_at,
				status,
				role
			FROM users
			WHERE login = $1
		`, login)
	return scanUser(row)
}
