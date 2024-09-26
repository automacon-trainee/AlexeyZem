package service

import (
	"context"
	"fmt"

	"project/internal/library/models"
)

type LibraryBD interface {
	GetAll(ctx context.Context) ([]models.Book, error)
	Create(ctx context.Context, book models.Book) error
	TakeBook(ctx context.Context, id int) (models.Book, error)
	ReturnBook(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (models.Book, error)
}

type LibraryServiceImpl struct {
	db LibraryBD
}

func NewLibraryServiceImpl(db LibraryBD) *LibraryServiceImpl {
	return &LibraryServiceImpl{
		db: db,
	}
}
func (s *LibraryServiceImpl) GetAll(ctx context.Context) ([]models.Book, error) {
	books, err := s.db.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all books: %w", err)
	}

	return books, nil
}

func (s *LibraryServiceImpl) Create(ctx context.Context, book models.Book) error {
	err := s.db.Create(ctx, book)
	if err != nil {
		return fmt.Errorf("create book: %w", err)
	}

	return nil
}

func (s *LibraryServiceImpl) Take(ctx context.Context, id int) (models.Book, error) {
	book, err := s.db.TakeBook(ctx, id)
	if err != nil {
		return models.Book{}, fmt.Errorf("get book: %w", err)
	}

	return book, nil
}

func (s *LibraryServiceImpl) Return(ctx context.Context, id int) error {
	err := s.db.ReturnBook(ctx, id)
	if err != nil {
		return fmt.Errorf("get book: %w", err)
	}

	return nil
}

func (s *LibraryServiceImpl) Get(ctx context.Context, id int) (models.Book, error) {
	book, err := s.db.GetByID(ctx, id)
	if err != nil {
		return models.Book{}, fmt.Errorf("get book: %w", err)
	}

	return book, nil
}