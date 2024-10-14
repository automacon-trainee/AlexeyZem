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
	createNewTable = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE 
	);`
	createUser = `INSERT INTO users (username, password, email)
        VALUES ($1, $2, $3)
        ON CONFLICT (email) DO NOTHING;`
	getUserByID    = `SELECT username, email FROM users WHERE id = $1`
	getList        = `SELECT username, email, password FROM users`
	getUserByEmail = `SELECT username, email, password FROM users WHERE email = $1`
)

var connect = make(map[string]*PostgressDataBase)

type PostgressDataBase struct {
	DB      *sql.DB
	metrics *metrics.ProxyMetrics
}

func StartPostgressDataBase(ctx context.Context, connStr string) (*PostgressDataBase, error) {
	if val, ok := connect[connStr]; ok {
		return val, nil
	}
	dataBase := &PostgressDataBase{}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return dataBase, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err = db.Ping(); err != nil {
		return dataBase, fmt.Errorf("failed to ping postgres: %w", err)
	}

	dataBase.DB = db
	dataBase.metrics = metrics.NewProxyMetrics()
	if err = dataBase.CreateNewUserTable(ctx); err != nil {
		return dataBase, fmt.Errorf("failed to create user table: %w", err)
	}
	connect[connStr] = dataBase

	return dataBase, err
}

func (db *PostgressDataBase) CreateNewUserTable(ctx context.Context) error {
	_, err := db.DB.ExecContext(ctx, createNewTable)

	return err
}

func (db *PostgressDataBase) Create(ctx context.Context, user models.User) error {
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

func (db *PostgressDataBase) GetByID(ctx context.Context, id string) (models.User, error) {
	metric := db.metrics.NewDurationHistogram("GetByID_method_histogram",
		"request GetById duration in second in DB",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()

	var user models.User

	row := db.DB.QueryRowContext(ctx, getUserByID, id)
	if err := row.Scan(&user.Username, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user with ID %s not found", id)
		}
		return user, err
	}

	return user, nil
}

func (db *PostgressDataBase) List(ctx context.Context) ([]models.User, error) {
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

	var (
		users []models.User
		user  models.User
	)

	for rows.Next() {
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

func (db *PostgressDataBase) GetByEmail(ctx context.Context, email string) (models.User, error) {
	metric := db.metrics.NewDurationHistogram("GetByEmail_method_histogram",
		"request GetByEmail duration in second in DB",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		metric.Observe(duration)
	}()

	var user models.User

	row := db.DB.QueryRowContext(ctx, getUserByEmail, email)
	if err := row.Scan(&user.Username, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user %s not found", email)
		}
		return user, err
	}

	return user, nil
}
