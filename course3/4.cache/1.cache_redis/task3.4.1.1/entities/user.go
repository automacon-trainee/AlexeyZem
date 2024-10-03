package entities

import (
	"encoding/json"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
