package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PrepareQuery(operation, table string, user User) (_ string, _ []any, _ error) {
	switch operation {
	case "insert":
		return squirrel.Insert(table).Columns("name", "email").Values(user.Name, user.Email).ToSql()
	case "update":
		return squirrel.Update(table).Set("name", user.Name).Set("email", user.Email).Where(squirrel.Eq{"id": user.ID}).ToSql()
	case "delete":
		return squirrel.Delete(table).Where(squirrel.Eq{"id": user.ID}).ToSql()
	case "select":
		return squirrel.Select("id", "name", "email").From(table).Where(squirrel.Eq{"id": user.ID}).ToSql()
	}
	return "", nil, errors.New("operation not supported")
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY,
    name TEXT,
    email TEXT
)`)
	if err != nil {
		return fmt.Errorf("error create table: %w", err)
	}

	return nil
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	queryStr, params, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return fmt.Errorf("error in prepare query: %w", err)
	}
	_, err = db.Exec(queryStr, params...)
	if err != nil {
		return fmt.Errorf("error in exec query: %w", err)
	}
	return nil
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return User{}, fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	queryStr, params, err := PrepareQuery("select", "users", User{ID: id})
	if err != nil {
		return User{}, fmt.Errorf("error in prepare query: %w", err)
	}
	rows, err := db.Query(queryStr, params...)
	if err != nil {
		return User{}, fmt.Errorf("error find user with id:%v: %w", id, err)
	}

	defer rows.Close()
	var name string
	var email string
	for rows.Next() {
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			return User{}, fmt.Errorf("error scan user: %w", err)
		}
	}
	return User{id, name, email}, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	queryStr, params, err := PrepareQuery("update", "users", user)
	if err != nil {
		return fmt.Errorf("error in prepare query: %w", err)
	}
	_, err = db.Exec(queryStr, params...)
	if err != nil {
		return fmt.Errorf("error update user: %w", err)
	}

	return nil
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	queryStr, params, err := PrepareQuery("delete", "users", User{ID: id})
	if err != nil {
		return fmt.Errorf("error in prepare query: %w", err)
	}
	_, err = db.Exec(queryStr, params...)
	if err != nil {
		return fmt.Errorf("error delete user: %w", err)
	}
	return nil
}

func main() {
	err := CreateUserTable()
	if err != nil {
		fmt.Printf("error in create table: %v\n", err)
	} else {
		fmt.Println("create table")
	}
	var user = User{Email: "test@example.com", Name: "john Doe"}
	err = InsertUser(user)
	if err != nil {
		fmt.Printf("error in insert user: %v\n", err)
	} else {
		fmt.Println("insert user")
	}
	user, err = SelectUser(3)
	if err != nil {
		fmt.Printf("error in select user: %v\n", err)
	} else {
		fmt.Println("select user:", user)
	}
	user.Email = "new@example.com"
	err = UpdateUser(user)
	if err != nil {
		fmt.Printf("error in update user: %v\n", err)
	} else {
		fmt.Println("update user:", user)
	}
	deletingID := 4
	err = DeleteUser(deletingID)
	if err != nil {
		fmt.Printf("error in delete user: %v\n", err)
	} else {
		fmt.Println("delete user")
	}
}
