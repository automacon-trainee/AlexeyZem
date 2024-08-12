package main

import (
	"strings"
	"testing"
)

type TestCaseCreateTable struct {
	user   User
	expect string
}

func TestCreateTable(t *testing.T) {
	sqlGenerator := &SQLLiteGenerator{}
	tests := []TestCaseCreateTable{
		{user: User{}, expect: `CREATE TABLE user (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		email VARCHAR(100) UNIQUE
	)`},
	}

	for _, test := range tests {
		res := sqlGenerator.CreateTableSQL(&test.user)
		if res != test.expect {
			t.Errorf("CreateTableSQL failed, expect: %s, actual: %s", test.expect, res)
		}
	}
}

type TestCaseInsertTable struct {
	user   User
	expect string
}

func TestInsertTable(t *testing.T) {
	sqlGenerator := &SQLLiteGenerator{}
	tests := []TestCaseInsertTable{
		{user: User{FirstName: "John", LastName: "Ken", ID: 19}, expect: "INSERT INTO user (id, first_name, last_name, email) VALUES"},
	}

	for _, test := range tests {
		res := sqlGenerator.CreateInsertSQL(test.user)
		if !strings.Contains(res, test.expect) {
			t.Errorf("CreateInsertSQL failed, expect: %s, actual: %s", test.expect, res)
		}
	}
}

type TestCaseTabler struct {
	user   User
	expect string
}

func TestTabler(t *testing.T) {
	tests := []TestCaseTabler{
		{user: User{}, expect: "user"},
	}
	for _, test := range tests {
		res := test.user.TableName()
		if res != test.expect {
			t.Errorf("TableName failed, expect: %s, actual: %s", test.expect, res)
		}
	}
}

func TestGenerateFakeUser(t *testing.T) {
	faker := GoFakeitGenerator{}
	user1 := faker.GenerateFakeUser()
	user2 := faker.GenerateFakeUser()
	user3 := faker.GenerateFakeUser()
	if user1 == user2 || user2 == user3 || user3 == user1 {
		t.Errorf("GenerateFakeUser failed, not random")
	}
}
