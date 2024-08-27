package repository

import (
	"database/sql"
	"fmt"

	"golibrary/entities"

	_ "github.com/lib/pq"
)

type LibraryRepo struct {
	db *sql.DB
}

type Librarer interface {
	CreateAuthor(author entities.Author) error
	CreateBook(book entities.Book) error
	CreateUser(user entities.User) error
	TakeBook(userID, bookID int) error
	ReturnBook(book entities.Book) error
	GetBookByID(id int) (entities.Book, error)
	GetAuthorByID(id int) (entities.Author, error)
	GetUserByID(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetAllBooksTakenByUser(userID int) ([]entities.Book, error)
	GetAllAuthors() ([]entities.Author, error)
	GetAllAuthorBooks(authorID int) ([]entities.Book, error)
	GetAllBooks() ([]entities.Book, error)
	HowManyAuthorsExist() (int, error)
	HowManyBooksExist() (int, error)
	HowManyUsersExist() (int, error)
}

func NewLibraryRepo(db *sql.DB) (Librarer, error) {
	service := &LibraryRepo{
		db: db,
	}

	err := service.CreateNewUserTable()
	if err != nil {
		return service, err
	}

	err = service.CreateNewBookTable()
	if err != nil {
		return service, err
	}

	err = service.CreateAuthorsTable()

	return service, err
}

func (ls *LibraryRepo) CreateNewUserTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE
	)`

	_, err := ls.db.Exec(newTableString)
	return err
}

func (ls *LibraryRepo) CreateNewBookTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		authorid INTEGER NOT NULL,
		usertakenid INTEGER DEFAULT 0,
		istaken BOOL DEFAULT false
	)`

	_, err := ls.db.Exec(newTableString)
	return err
}

func (ls *LibraryRepo) CreateAuthorsTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL UNIQUE
	)`

	_, err := ls.db.Exec(newTableString)
	return err
}

func (ls *LibraryRepo) CreateAuthor(author entities.Author) error {
	tx, err := ls.db.Begin()
	if err != nil {
		return err
	}

	query := ` INSERT INTO authors (name) VALUES ($1)`

	result, err := tx.Exec(query, author.Name)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with username %s already exists", author.Name)
	}

	for _, book := range author.Books {
		book.AuthorID = author.ID
		err = ls.CreateBook(book)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	
	return tx.Commit()
}

func (ls *LibraryRepo) CreateBook(book entities.Book) error {
	query := `SELECT name FROM authors WHERE id = $1`
	row, err := ls.db.Exec(query, book.AuthorID)
	if err != nil {
		return err
	}
	rowCount, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowCount == 0 {
		return fmt.Errorf("author with id %d doesnt exist", book.AuthorID)
	}

	tx, err := ls.db.Begin()
	if err != nil {
		return err
	}

	query = `
		INSERT INTO books (name, authorid)
		VALUES ($1, $2)
	`

	_, err = tx.Exec(query, book.Name, book.AuthorID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (ls *LibraryRepo) CreateUser(user entities.User) error {
	tx, err := ls.db.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (username)
		VALUES ($1)
	`

	_, err = tx.Exec(query, user.Username)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (ls *LibraryRepo) TakeBook(userID, bookID int) error {
	query := `SELECT istaken FROM books WHERE id = $1`
	var bookIsTaken bool
	rows, err := ls.db.Query(query, bookID)
	if err != nil {
		return err
	}

	_ = rows.Scan(&bookIsTaken)
	if bookIsTaken {
		return fmt.Errorf("book with id %d is already taken", bookID)
	}

	query = `UPDATE books SET usertakenid = $1, istaken = $2 WHERE id = $3`
	_, err = ls.db.Exec(query, userID, true, bookID)

	return err
}

func (ls *LibraryRepo) ReturnBook(book entities.Book) error {
	query := `SELECT istaken FROM books WHERE id = $1`
	var bookIsTaken bool
	rows, err := ls.db.Query(query, book.ID)
	if err != nil {
		return err
	}

	_ = rows.Scan(&bookIsTaken)
	if !bookIsTaken {
		return fmt.Errorf("book with id %d not taken", book.ID)
	}

	query = `UPDATE books SET istaken = $1 WHERE id = $2`
	_, err = ls.db.Exec(query, false, book.ID)

	return err
}

func (ls *LibraryRepo) GetBookByID(id int) (entities.Book, error) {
	query := `SELECT (name, authorid, usertakenid) FROM books WHERE id = $1`
	row := ls.db.QueryRow(query, id)
	book := entities.Book{ID: id}
	err := row.Scan(&book.Name, &book.AuthorID, &book.TakenBy)
	if err != nil {
		return book, err
	}

	author, err := ls.GetAuthorByID(book.AuthorID)
	if err != nil {
		return book, err
	}

	book.Author = author.Name
	return book, nil
}

func (ls *LibraryRepo) GetAuthorByID(id int) (entities.Author, error) {
	query := `SELECT name FROM authors WHERE id = $1`
	row := ls.db.QueryRow(query, id)
	var author entities.Author
	err := row.Scan(&author.Name)
	if err != nil {
		return author, err
	}
	books, err := ls.GetAllAuthorBooks(author.ID)
	if err != nil {
		return author, err
	}
	author.ID = id
	author.Books = books
	return author, nil
}

func (ls *LibraryRepo) GetUserByID(id int) (entities.User, error) {
	query := `SELECT username FROM users WHERE id = $1`
	row := ls.db.QueryRow(query, id)
	var user entities.User
	err := row.Scan(&user.Username)
	if err != nil {
		return user, err
	}
	books, err := ls.GetAllBooksTakenByUser(id)
	if err != nil {
		return user, err
	}
	user.BooksTaken = books
	user.ID = id
	return user, nil
}

func (ls *LibraryRepo) GetAllBooksTakenByUser(userID int) ([]entities.Book, error) {
	query := `SELECT (id, name, authorid) FROM books WHERE usertakenid = $1 AND istaken = $2`
	rows, err := ls.db.Query(query, userID, true)
	var books []entities.Book
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book entities.Book
		err = rows.Scan(&book.ID, &book.Name, &book.AuthorID)
		if err != nil {
			return nil, err
		}
		book.TakenBy = userID
		author, err := ls.GetAuthorByID(book.AuthorID)
		if err != nil {
			return books, err
		}
		book.Author = author.Name
		books = append(books, book)
	}

	return books, nil
}

func (ls *LibraryRepo) GetAllUsers() ([]entities.User, error) {
	query := `SELECT * FROM users`
	rows, err := ls.db.Query(query)
	var users []entities.User
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}

		books, err := ls.GetAllBooksTakenByUser(user.ID)
		if err != nil {
			return nil, err
		}

		user.BooksTaken = books
		users = append(users, user)
	}

	return users, nil
}

func (ls *LibraryRepo) GetAllAuthorBooks(authorID int) ([]entities.Book, error) {
	query := `SELECT (id, name, usertakenid)  FROM books WHERE authorid = $1`
	rows, err := ls.db.Query(query, authorID)
	var books []entities.Book
	if err != nil {
		return books, err
	}
	author, err := ls.GetAuthorByID(authorID)
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book entities.Book
		err = rows.Scan(&book.ID, &book.Name, &book.TakenBy)

		if err != nil {
			return nil, err
		}

		book.Author = author.Name
		books = append(books, book)
	}

	return books, nil
}

func (ls *LibraryRepo) GetAllAuthors() ([]entities.Author, error) {
	query := `SELECT id, name FROM authors`
	rows, err := ls.db.Query(query)
	var authors []entities.Author
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author entities.Author
		err = rows.Scan(&author.ID, &author.Name)

		if err != nil {
			return nil, err
		}

		authorBooks, err := ls.GetAllAuthorBooks(author.ID)
		if err != nil {
			return nil, err
		}

		author.Books = authorBooks
		authors = append(authors, author)
	}

	return authors, nil
}

func (ls *LibraryRepo) GetAllBooks() ([]entities.Book, error) {
	query := `SELECT (id, name, authorid, usernakenid) FROM books`
	rows, err := ls.db.Query(query)
	var books []entities.Book
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book entities.Book
		_ = rows.Scan(&book.ID, &book.Name, &book.AuthorID, &book.TakenBy)

		author, err := ls.GetAuthorByID(book.AuthorID)
		if err != nil {
			return books, err
		}

		book.Author = author.Name
		books = append(books, book)
	}

	return books, nil
}

func (ls *LibraryRepo) HowManyAuthorsExist() (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM authors`
	var authorsQuantity int
	err := ls.db.QueryRow(query).Scan(&authorsQuantity)
	if err != nil {
		return 0, err
	}

	return authorsQuantity, nil
}

func (ls *LibraryRepo) HowManyBooksExist() (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM books`
	var booksQuantity int
	err := ls.db.QueryRow(query).Scan(&booksQuantity)
	if err != nil {
		return 0, err
	}

	return booksQuantity, nil
}

func (ls *LibraryRepo) HowManyUsersExist() (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM users`
	var usersQuantity int
	err := ls.db.QueryRow(query).Scan(&usersQuantity)
	if err != nil {
		return 0, err
	}

	return usersQuantity, nil
}
