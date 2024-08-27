package service

import (
	"math/rand"
	"time"

	"golibrary/entities"
	"golibrary/internal/repository"

	"github.com/brianvoe/gofakeit"
)

type LibraryFacade struct {
	Repo repository.Librarer
}

type Servicer interface {
	StartService() error
	TakeBook(userID, bookID int) error
	ReturnBook(book entities.Book) error
	AllUsersInfo() ([]entities.User, error)
	AllAuthorsInfo() ([]entities.Author, error)
	AddBook(book entities.Book) error
	AddUser(user entities.User) error
	AddAuthor(author entities.Author) error
	GetAllBooks() ([]entities.Book, error)
	BookInfo(bookID int) (entities.Book, error)
	AuthorInfo(authorID int) (entities.Author, error)
	UserInfo(userID int) (entities.User, error)
}

func NewLibraryFacade(repo repository.Librarer) (*LibraryFacade, error) {
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
