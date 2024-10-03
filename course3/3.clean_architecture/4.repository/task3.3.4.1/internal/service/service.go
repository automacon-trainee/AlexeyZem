package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"projectrepo/internal"
	"projectrepo/internal/controller"
	"projectrepo/internal/models"
)

type Conditions struct {
	Limit  int
	Offset int
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User, username string) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Delete(ctx context.Context, username string) error
	List(ctx context.Context, c Conditions) ([]*models.User, error)
	CreateNewTable() error
}

type ImplService struct {
	repo UserRepository
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
	user, err = i.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return
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
	users, err = i.repo.List(ctx, Conditions{Limit: lim, Offset: off})
	return
}

func (i *ImplService) Update(username string, user *models.User) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = i.repo.Update(ctx, user, username)
	return
}

func NewService(repo UserRepository) controller.UserService {
	return &ImplService{repo: repo}
}
