package service

import (
	"context"

	"project/internal/API/gRPCBook"
	lib "project/internal/library/models"
	"project/internal/profile/models"
)

type ProfileRepo interface {
	Create(ctx context.Context, profile models.Profile) error
	GetProfile(ctx context.Context, id int) (*models.Profile, []int, error)
	TakeBook(ctx context.Context, profileID, bookID int) error
	ReturnBook(ctx context.Context, profileID, bookID int) error
}

type ProfileServiceImpl struct {
	repo    ProfileRepo
	libgRPC gRPCBook.BookServiceClient
}

func NewProfileServiceImpl(repo ProfileRepo, lib gRPCBook.BookServiceClient) *ProfileServiceImpl {
	return &ProfileServiceImpl{
		repo:    repo,
		libgRPC: lib,
	}
}

func (s *ProfileServiceImpl) CreateProfile(ctx context.Context, profile models.Profile) error {
	err := s.repo.Create(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProfileServiceImpl) GetProfile(ctx context.Context, id int) (*models.Profile, error) {
	profile, booksID, err := s.repo.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}

	var books []lib.Book
	for _, id := range booksID {
		book, err := s.libgRPC.Get(ctx, &gRPCBook.ID{Id: int64(id)})
		if err != nil {
			return nil, err
		}
		books = append(books, lib.Book{ID: id, Title: book.Title, Author: book.Author})
	}

	profile.TakenBook = books

	return profile, nil
}

func (s *ProfileServiceImpl) TakeBook(ctx context.Context, profileID, bookID int) error {
	_, err := s.libgRPC.Take(ctx, &gRPCBook.ID{Id: int64(bookID)})
	if err != nil {
		return err
	}

	err = s.repo.TakeBook(ctx, profileID, bookID)
	if err != nil {
		return err
	}

	return nil

}

func (s *ProfileServiceImpl) ReturnBook(ctx context.Context, profileID, bookID int) error {
	_, err := s.libgRPC.Return(ctx, &gRPCBook.ID{Id: int64(bookID)})
	if err != nil {
		return err
	}

	err = s.repo.ReturnBook(ctx, profileID, bookID)
	if err != nil {
		return err
	}

	return nil
}
