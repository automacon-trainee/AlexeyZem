package service

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit"
	"golibrary/entities"
)

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

type LibraryFacade struct {
	Repo Librarer
}

func NewLibraryFacade(repo Librarer) (*LibraryFacade, error) {
	lf := &LibraryFacade{Repo: repo}
	err := lf.StartService()
	return lf, err
}

func (lf *LibraryFacade) StartService() error {
	authorsNumber, err := lf.Repo.HowManyAuthorsExist()
	if err != nil {
		return err
	}
	var (
		minAuthors     = 10
		minBooksNumber = 100
		minUsersNumber = 50
	)
	randomer := rand.New(rand.NewSource(time.Now().UnixNano()))
	if authorsNumber < minAuthors {
		for i := 0; i < (10 - authorsNumber); i++ {
			authorID := randomer.Intn(authorsNumber) + 1
			err = lf.Repo.CreateAuthor(entities.Author{
				ID:   authorID,
				Name: gofakeit.Name(),
				Books: []entities.Book{
					{
						Name:     gofakeit.Word(),
						AuthorID: authorID,
					},
					{
						Name:     gofakeit.Word(),
						AuthorID: authorID,
					},
				},
			})
			if err != nil {
				return err
			}
		}
		authorsNumber = minAuthors
	}

	booksNumber, err := lf.Repo.HowManyBooksExist()
	if err != nil {
		return err
	}

	if booksNumber < minBooksNumber {
		for i := 0; i < (100 - booksNumber); i++ {
			authorID := randomer.Intn(authorsNumber) + 1

			book := entities.Book{
				Name:     gofakeit.Word(),
				AuthorID: authorID,
			}

			err = lf.Repo.CreateBook(book)
			if err != nil {
				return err
			}
		}
	}

	usersNumber, err := lf.Repo.HowManyUsersExist()
	if err != nil {
		return err
	}

	if usersNumber < minUsersNumber {
		for i := 0; i < (50 - usersNumber); i++ {
			user := entities.User{
				Username: gofakeit.Name(),
			}
			err := lf.Repo.CreateUser(user)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (lf *LibraryFacade) TakeBook(userID, bookID int) error {
	return lf.Repo.TakeBook(userID, bookID)
}

func (lf *LibraryFacade) ReturnBook(book entities.Book) error {
	return lf.Repo.ReturnBook(book)
}

func (lf *LibraryFacade) AllUsersInfo() ([]entities.User, error) {
	return lf.Repo.GetAllUsers()
}

func (lf *LibraryFacade) AllAuthorsInfo() ([]entities.Author, error) {
	return lf.Repo.GetAllAuthors()
}

func (lf *LibraryFacade) AddBook(book entities.Book) error {
	return lf.Repo.CreateBook(book)
}

func (lf *LibraryFacade) AddUser(user entities.User) error {
	return lf.Repo.CreateUser(user)
}

func (lf *LibraryFacade) AddAuthor(author entities.Author) error {
	return lf.Repo.CreateAuthor(author)
}

func (lf *LibraryFacade) GetAllBooks() ([]entities.Book, error) {
	return lf.Repo.GetAllBooks()
}

func (lf *LibraryFacade) BookInfo(bookID int) (entities.Book, error) {
	return lf.Repo.GetBookByID(bookID)
}

func (lf *LibraryFacade) AuthorInfo(authorID int) (entities.Author, error) {
	return lf.Repo.GetAuthorByID(authorID)
}

func (lf *LibraryFacade) UserInfo(userID int) (entities.User, error) {
	return lf.Repo.GetUserByID(userID)
}
