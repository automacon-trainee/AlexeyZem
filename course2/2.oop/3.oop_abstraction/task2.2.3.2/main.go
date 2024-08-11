package main

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	panic("implement me")
}

type Tabler interface {
	TableName() string
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLLiteGenerator struct {
	ID int
}

func (s *SQLLiteGenerator) CreateTableSQL(_ Tabler) string {
	panic("implement me")
}

func (s *SQLLiteGenerator) CreateInsertSQL(_ Tabler) string {
	panic("implement me")
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct {
	ID int
}

func (f *GoFakeitGenerator) GenerateFakeUser() User {
	panic("implement me")
}

func main() {}
