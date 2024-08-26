package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"projectrepo/internal"
	"projectrepo/internal/models"
	"projectrepo/repository"
)

type UserService interface {
	Create(user *models.User) (err error)
	Delete(username string) (err error)
	Get(username string) (user *models.User, err error)
	GetAll(limit, offset string) (users []*models.User, err error)
	Update(username string, user *models.User) (err error)
}

type ImplService struct {
	repo repository.UserRepository
}

func (i *ImplService) Create(user *models.User) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = i.repo.Create(ctx, user)
	return
}

func (i *ImplService) Delete(username string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = i.repo.Delete(ctx, username)
	return
}

func (i *ImplService) Get(username string) (user *models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	userRepo, err := i.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return userRepo.ToModelsUser(), nil
}

func (i *ImplService) GetAll(limit, offset string) (users []*models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if limit == "" {
		limit = "1"
	}
	if offset == "" {
		offset = "0"
	}
	lim, err := strconv.Atoi(limit)
	if err != nil {
		return nil, fmt.Errorf("%w:%w", err, internal.BadRequestError)
	}
	off, err := strconv.Atoi(offset)
	if err != nil {
		return nil, fmt.Errorf("%w:%w", err, internal.BadRequestError)
	}
	users, err = i.repo.List(ctx, repository.Conditions{Limit: lim, Offset: off})
	return
}

func (i *ImplService) Update(username string, user *models.User) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = i.repo.Update(ctx, user, username)
	return
}

func NewService(repo repository.UserRepository) UserService {
	return &ImplService{repo: repo}
}
