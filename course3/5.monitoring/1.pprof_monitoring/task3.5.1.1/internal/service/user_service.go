package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/jwtauth"

	"github.com/go-redis/redis"

	"pprof/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByID(ctx context.Context, id string) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
}

type UserServiceImpl struct {
	repo  UserRepository
	token *jwtauth.JWTAuth
}

func NewUserServiceImpl(repo UserRepository, token *jwtauth.JWTAuth) *UserServiceImpl {
	return &UserServiceImpl{
		repo:  repo,
		token: token,
	}
}

func (s *UserServiceImpl) CreateUser(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.repo.Create(context.Background(), user)
}

func (s *UserServiceImpl) AuthUser(user models.User) (string, error) {
	userBD, err := s.repo.GetByEmail(context.Background(), user.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userBD.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("wrong password")
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}
	_, token, _ := s.token.Encode(claims)
	return token, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.repo.List(context.Background())
}

func (s *UserServiceImpl) GetUserByEmail(email string) (models.User, error) {
	return s.repo.GetByEmail(context.Background(), email)
}

type UserServiceProxy struct {
	userService *UserServiceImpl
	client      *redis.Client
}

func NewUserServiceProxy(userService *UserServiceImpl, client *redis.Client) *UserServiceProxy {
	return &UserServiceProxy{
		userService: userService,
		client:      client,
	}
}

func (s *UserServiceProxy) CreateUser(user models.User) error {
	return s.userService.CreateUser(user)
}

func (s *UserServiceProxy) AuthUser(user models.User) (string, error) {
	return s.userService.AuthUser(user)
}

func (s *UserServiceProxy) GetAllUsers() ([]models.User, error) {
	data, err := s.client.Get("allUsers").Result()
	if err != nil {
		users, errBD := s.userService.GetAllUsers()
		if errBD != nil {
			return nil, errBD
		}
		if errors.Is(err, redis.Nil) {
			s.client.Set("allUsers", users, time.Minute*5)
		}
		return users, nil
	}
	var users []models.User
	err = json.Unmarshal([]byte(data), &users)
	return users, err
}

func (s *UserServiceProxy) GetUserByEmail(email string) (models.User, error) {
	data, err := s.client.Get(email).Result()
	if err != nil {
		user, errBD := s.userService.GetUserByEmail(email)
		if errBD != nil {
			return models.User{}, err
		}
		if errors.Is(err, redis.Nil) {
			s.client.Set(email, user, time.Hour)
		}
		return user, nil
	}
	user := models.User{}
	err = json.Unmarshal([]byte(data), &user)
	return user, err
}
