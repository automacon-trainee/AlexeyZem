package entities

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	AuthorID int    `json:"authorid"`
	TakenBy  int    `json:"takenby"`
}

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	BooksTaken []Book `json:"bookstaken"`
}

type Author struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}
