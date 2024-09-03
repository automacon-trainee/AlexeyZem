package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/jwtauth"

	"github.com/go-redis/redis"

	"metrics/internal/metrics"
	"metrics/internal/models"
	"metrics/internal/repository"
)

type UserService interface {
	CreateUser(user models.User) error
	AuthUser(user models.User) (string, error)
	GetUserByEmail(email string) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserServiceImpl struct {
	repo  repository.UserRepository
	token *jwtauth.JWTAuth
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

func NewUserServiceImpl(repo repository.UserRepository, token *jwtauth.JWTAuth) UserService {
	return &UserServiceImpl{
		repo:  repo,
		token: token,
	}
}

type UserServiceProxy struct {
	userService UserService
	client      *redis.Client
	metrics     *metrics.ProxyMetrics
}

func (s *UserServiceProxy) CreateUser(user models.User) error {
	return s.userService.CreateUser(user)
}

func (s *UserServiceProxy) AuthUser(user models.User) (string, error) {
	return s.userService.AuthUser(user)
}

func (s *UserServiceProxy) GetAllUsers() ([]models.User, error) {
	histogram := s.metrics.NewDurationHistogram("GetAllUser_cache_histogram", "time for cache getAllUser",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()

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

	duration := time.Since(start).Seconds()
	histogram.Observe(duration)
	var users []models.User
	err = json.Unmarshal([]byte(data), &users)
	return users, err
}

func (s *UserServiceProxy) GetUserByEmail(email string) (models.User, error) {
	histogram := s.metrics.NewDurationHistogram("GetUserByEmail_endpoint_histogram", "time for cache getUserByEmail",
		prometheus.LinearBuckets(0.1, 0.1, 10))
	start := time.Now()

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

	duration := time.Since(start).Seconds()
	histogram.Observe(duration)
	user := models.User{}
	err = json.Unmarshal([]byte(data), &user)
	return user, err
}

func NewUserServiceProxy(userService UserService, client *redis.Client) UserService {
	return &UserServiceProxy{
		userService: userService,
		client:      client,
		metrics:     metrics.NewProxyMetrics(),
	}
}
