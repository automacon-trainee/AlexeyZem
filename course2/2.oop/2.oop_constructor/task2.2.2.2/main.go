package main

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func NewUser(id int, opts ...UserOption) *User {
	u := &User{
		ID: id,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

func WithUserName(name string) UserOption {
	return func(u *User) {
		u.Username = name
	}
}

func WithUserEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithUserRole(role string) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func main() {}
