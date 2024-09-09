package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"projectrepo/internal"
	"projectrepo/internal/models"
	"projectrepo/internal/service"
)

type DB interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

type UserRepositoryImplSQL struct {
	db DB
}

func NewUserRepository(db DB) service.UserRepository {
	return &UserRepositoryImplSQL{
		db: db,
	}
}

func (u *UserRepositoryImplSQL) Create(ctx context.Context, user *models.User) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if u.isDeleted(ctx, user.Username) {
		return u.updateUser(ctx, user, user.Username)
	}
	res, err := tx.ExecContext(ctx, `INSERT INTO users (username, password) VALUES($1, $2)`, user.Username, user.Password)
	if err != nil {
		_ = tx.Rollback()
		log.Println(err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		log.Println(err)
		return err
	}
	if rows == 0 {
		_ = tx.Rollback()
		log.Println(err)
		return internal.BadRequestError
	}
	err = tx.Commit()
	return err
}

func (u *UserRepositoryImplSQL) Update(ctx context.Context, user *models.User, username string) error {
	if u.isDeleted(ctx, user.Username) {
		return internal.BadRequestError
	}
	return u.updateUser(ctx, user, username)
}

func (u *UserRepositoryImplSQL) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT * FROM users WHERE username = $1`
	rows, err := u.db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, internal.BadRequestError
	}
	var user User
	err = rows.Scan(&user.Username, &user.Password, &user.IsExist)
	if !user.IsExist {
		return nil, internal.DeletedError
	}
	res := user.ToModelsUser()
	return res, err
}

func (u *UserRepositoryImplSQL) Delete(ctx context.Context, username string) error {
	query := `UPDATE users SET isExist = $1 WHERE username = $2`
	if u.isDeleted(ctx, username) {
		return internal.BadRequestError
	}
	res, err := u.db.ExecContext(ctx, query, false, username)
	if err != nil {
		log.Println("error deleting user Exec:", err)
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return internal.BadRequestError
	}
	return err
}

func (u *UserRepositoryImplSQL) List(ctx context.Context, c service.Conditions) ([]*models.User, error) {
	query := `SELECT * FROM users limit $1 offset $2`
	rows, err := u.db.QueryContext(ctx, query, c.Limit, c.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Username, &user.Password, &user.IsExist)
		if err != nil {
			return nil, err
		}
		if user.IsExist {
			users = append(users, user.ToModelsUser())
		}
	}
	return users, nil
}

func (u *UserRepositoryImplSQL) CreateNewTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    isExist boolean default true
    );`
	_, err := u.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImplSQL) isDeleted(ctx context.Context, username string) bool {
	userMy, err := u.GetByUsername(ctx, username)
	if userMy == nil && errors.Is(err, internal.DeletedError) {
		return true
	}
	return false
}

func (u *UserRepositoryImplSQL) updateUser(ctx context.Context, user *models.User, username string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := `UPDATE users set username = $1, password = $2, isExist = $3 WHERE username = $4`
	res, err := tx.ExecContext(ctx, query, user.Username, user.Password, true, username)
	if err != nil {
		_ = tx.Rollback()
		log.Println(err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		log.Println(err)
		return err
	}
	if rows == 0 {
		_ = tx.Rollback()
		log.Println(err)
		return internal.BadRequestError
	}
	err = tx.Commit()
	log.Println(err)
	return err
}
