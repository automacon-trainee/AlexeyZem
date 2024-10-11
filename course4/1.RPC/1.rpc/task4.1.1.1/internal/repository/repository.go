package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

const (
	createTable = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE 
	);`
	createUser = `
        INSERT INTO users (username, password, email)
        VALUES ($1, $2, $3)
        ON CONFLICT (email) DO NOTHING;`
	getUser    = `SELECT username, email FROM users WHERE id = $1`
	getList    = `SELECT username, email, password FROM users`
	getByEmail = `SELECT username, email, password FROM users WHERE email = $1`
)

type PostgresDataBase struct {
	DB      *sql.DB
	metrics *metrics.ProxyMetrics
}

func StartPostgressDataBase(ctx context.Context, connStr string) (*PostgresDataBase, error) {
	var dataBase PostgresDataBase

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &dataBase, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return &dataBase, fmt.Errorf("failed to ping postgres: %w", err)
	}

	dataBase.DB = db
	dataBase.metrics = metrics.NewProxyMetrics()
	err = dataBase.CreateNewUserTable(ctx)

	return &dataBase, err
}

func (db *PostgresDataBase) CreateNewUserTable(ctx context.Context) error {
	_, err := db.DB.ExecContext(ctx, createTable)

	return err
}

func (db *PostgresDataBase) Create(ctx context.Context, user models.User) error {
	metric := db.metrics.NewDurationHistogram("Create_method_histogram",
		"request Create duration in second in DB",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()

	result, err := db.DB.ExecContext(ctx, createUser, user.Username, user.Password, user.Email)
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

func (db *PostgresDataBase) GetByID(ctx context.Context, id string) (models.User, error) {
	metric := db.metrics.NewDurationHistogram("GetByID_method_histogram",
		"request GetById duration in second in DB",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()

	var user models.User

	row := db.DB.QueryRowContext(ctx, getUser, id)
	if err := row.Scan(&user.Username, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user with ID %s not found", id)
		}

		return user, err
	}

	return user, nil
}

func (db *PostgresDataBase) List(ctx context.Context) ([]models.User, error) {
	metric := db.metrics.NewDurationHistogram("List_method_histogram",
		"request List duration in second in DB",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()

	rows, err := db.DB.QueryContext(ctx, getList)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Username, &user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *PostgresDataBase) GetByEmail(ctx context.Context, email string) (models.User, error) {
	metric := db.metrics.NewDurationHistogram("GetByEmail_method_histogram",
		"request GetByEmail duration in second in DB",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()

	var user models.User

	row := db.DB.QueryRowContext(ctx, getByEmail, email)
	if err := row.Scan(&user.Username, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user %s not found", email)
		}
	}

	return user, nil
}
