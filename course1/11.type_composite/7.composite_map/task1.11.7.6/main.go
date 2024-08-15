package main

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUser(users []User) []User {
	m := make(map[string]struct{})
	for i, user := range users {
		if _, ok := m[user.Nickname]; !ok {
			m[user.Nickname] = struct{}{}
		} else {
			users = append(users[:i], users[i+1:]...)
		}
	}
	return users
}

func main() {}
