package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
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
    age INTEGER
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
	_, err = db.Exec(`INSERT INTO users(name, age) VALUES (?, ?)`, user.Name, user.Age)

	if err != nil {
		return fmt.Errorf("error insert user: %w", err)
	}

	return nil
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return User{}, fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		return User{}, fmt.Errorf("error find user with id:%v: %w", id, err)
	}

	defer rows.Close()
	var name string
	var age int
	for rows.Next() {
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			return User{}, fmt.Errorf("error scan user: %w", err)
		}
	}
	return User{id, name, age}, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return fmt.Errorf("error in open: %w", err)
	}
	defer db.Close()
	_, err = db.Exec(`UPDATE users SET name = ?, age = ? WHERE id = ?`, user.Name, user.Age, user.ID)
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
	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, id)
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
	var user = User{Age: 31, Name: "Last User"}
	err = InsertUser(user)
	if err != nil {
		fmt.Printf("error in insert user: %v\n", err)
	} else {
		fmt.Println("insert user")
	}
	user, err = SelectUser(2)
	if err != nil {
		fmt.Printf("error in select user: %v\n", err)
	} else {
		fmt.Println("select user:", user)
	}
	user.Age = 30
	err = UpdateUser(user)
	if err != nil {
		fmt.Printf("error in update user: %v\n", err)
	} else {
		fmt.Println("update user:", user)
	}
	deletingID := 6
	err = DeleteUser(deletingID)
	if err != nil {
		fmt.Printf("error in delete user: %v\n", err)
	} else {
		fmt.Println("delete user")
	}
}
