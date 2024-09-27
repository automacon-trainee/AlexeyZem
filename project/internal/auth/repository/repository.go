package repository

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"

	"project/internal/auth/models"
	"project/internal/myerror"
)

type PostgressDataBase struct {
	db *sql.DB
}

func NewPostgresDataBase(db *sql.DB) *PostgressDataBase {
	return &PostgressDataBase{
		db: db,
	}
}

func (db *PostgressDataBase) CreateNewUserTable(ctx context.Context) error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE 
	);`

	_, err := db.db.ExecContext(ctx, newTableString)
	return err
}

func (db *PostgressDataBase) Create(ctx context.Context, user *models.User) error {
	query := `
        INSERT INTO users (username, password, email)
        VALUES ($1, $2, $3)
        ON CONFLICT (email) DO NOTHING;
    `

	result, err := db.db.ExecContext(ctx, query, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return myerror.ErrUserAlreadyExists
	}

	return nil
}

func (db *PostgressDataBase) GetByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	query := `SELECT username, email FROM users WHERE id = $1`

	row := db.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Username, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, myerror.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (db *PostgressDataBase) List(ctx context.Context) ([]models.User, error) {
	query := `SELECT username, email, password FROM users`
	rows, err := db.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Username, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *PostgressDataBase) GetByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT username, email, password FROM users WHERE email = $1`
	var user models.User

	row := db.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.Username, &user.Email, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, myerror.ErrUserNotFound
		}
	}

	return user, err
}
