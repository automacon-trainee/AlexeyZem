package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

const userTableName = "user"

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return userTableName
}

type Tabler interface {
	TableName() string
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(user User) string
}

type SQLLiteGenerator struct {
	ID int
}

func (s *SQLLiteGenerator) CreateTableSQL(user Tabler) string {
	tableName := user.TableName()
	res := fmt.Sprintf(`CREATE TABLE %s (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	)`, tableName)

	return res
}

func (s *SQLLiteGenerator) CreateInsertSQL(user User) string {
	return fmt.Sprintf("INSERT INTO %s (id, first_name, last_name, email) VALUES (%d, %s, %s, %s)",
		user.TableName(), user.ID, user.FirstName, user.LastName, user.Email)
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct {
	ID int
}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	user := User{}
	g.ID++
	user.ID = g.ID
	user.FirstName = gofakeit.FirstName()
	user.LastName = gofakeit.LastName()
	user.Email = gofakeit.Email()
	return user
}

func main() {}
