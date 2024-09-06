// Оптимизируй использование внешней библиотеки в программе на языке Golang.
// Присутствуют 3 реализации интерфейса MarshalUnmarshaler, которые будут использовать
// стандартную библиотеку, библиотеку jsoniter и библиотеку easyjson.
// Присутствует бенчмарк, который будет сравнивать производительность этих реализаций.
package main

import (
	"encoding/json"
	"fmt"

	jsointer "github.com/json-iterator/go"
	"github.com/mailru/easyjson"

	"jsointer/model"
)

type MarshalUnMarshaler interface {
	Marshal(user *model.User) ([]byte, error)
	Unmarshal([]byte) (model.User, error)
}

type StandartJson struct{}

func (s *StandartJson) Marshal(user *model.User) ([]byte, error) {
	return json.Marshal(user)
}

func (s *StandartJson) Unmarshal(data []byte) (model.User, error) {
	var user model.User
	err := json.Unmarshal(data, &user)
	return user, err
}

type EasyJson struct{}

func (s *EasyJson) Marshal(user *model.User) ([]byte, error) {
	return easyjson.Marshal(user)
}

func (s *EasyJson) Unmarshal(data []byte) (model.User, error) {
	var user model.User
	err := easyjson.Unmarshal(data, &user)
	return user, err
}

type Jsointer struct{}

func (s *Jsointer) Marshal(user *model.User) ([]byte, error) {
	my := jsointer.ConfigCompatibleWithStandardLibrary
	return my.Marshal(user)
}

func (s *Jsointer) Unmarshal(data []byte) (model.User, error) {
	var user model.User
	my := jsointer.ConfigCompatibleWithStandardLibrary
	err := my.Unmarshal(data, &user)
	return user, err
}

func GenerateUser(count int) []model.User {
	users := make([]model.User, count)
	for i := 0; i < count; i++ {
		users[i] = model.User{
			ID:       i,
			Username: fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("pass%d", i),
			Age:      i,
			Email:    fmt.Sprintf("email%d", i),
		}
	}
	return users
}

func main() {
	users := GenerateUser(5)
	stJS := &StandartJson{}
	eaJS := &EasyJson{}
	js := &Jsointer{}
	for _, user := range users {
		data, err := stJS.Marshal(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
		data, err = eaJS.Marshal(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
		data, err = js.Marshal(&user)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
