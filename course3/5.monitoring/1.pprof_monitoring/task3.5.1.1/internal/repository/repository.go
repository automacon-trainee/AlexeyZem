package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"pprof/internal/models"
)

type PostgressDataBase struct {
	DB *sql.DB
}

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id string) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
}

func StartPostgressDataBase(ctx context.Context, connStr string) (*PostgressDataBase, error) {
	dataBase := &PostgressDataBase{}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return dataBase, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return dataBase, fmt.Errorf("failed to ping postgres: %w", err)
	}

	dataBase.DB = db
	err = dataBase.CreateNewUserTable(ctx)
	return dataBase, err
}

func (db *PostgressDataBase) CreateNewUserTable(ctx context.Context) error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE 
	);`

	_, err := db.DB.ExecContext(ctx, newTableString)
	return err
}

func (db *PostgressDataBase) Create(ctx context.Context, user models.User) error {
	query := `
        INSERT INTO users (username, password, email)
        VALUES ($1, $2, $3)
        ON CONFLICT (email) DO NOTHING;
    `

	result, err := db.DB.ExecContext(ctx, query, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	return nil
}

func (db *PostgressDataBase) GetByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	query := `SELECT username, email FROM users WHERE id = $1`

	row := db.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.Username, &user.Email)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user with ID %s not found", id)
		}

		return user, err
	}

	return user, nil
}

func (db *PostgressDataBase) List(ctx context.Context) ([]models.User, error) {
	query := `SELECT username, email FROM users`
	rows, err := db.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Username, &user.Email)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *PostgressDataBase) GetByEmail(ctx context.Context, email string) (models.User, error) {
	query := `SELECT username, email FROM users WHERE email = $1`
	var user models.User

	row := db.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.Username, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user %s not found", email)
		}
	}

	return user, err
}
