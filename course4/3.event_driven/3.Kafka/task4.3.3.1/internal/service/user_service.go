package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-redis/redis"

	"metrics/internal/metrics"
	"metrics/internal/models"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserServiceImpl(repo UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.repo.List(context.Background())
}

func (s *UserServiceImpl) GetUserByEmail(email string) (models.User, error) {
	return s.repo.GetByEmail(context.Background(), email)
}

type UserService interface {
	GetUserByEmail(email string) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserServiceProxy struct {
	userService UserService
	client      *redis.Client
	metrics     *metrics.ProxyMetrics
}

func NewUserServiceProxy(userService UserService, client *redis.Client) *UserServiceProxy {
	return &UserServiceProxy{
		userService: userService,
		client:      client,
		metrics:     metrics.NewProxyMetrics(),
	}
}

func (s *UserServiceProxy) GetAllUsers() ([]models.User, error) {
	histogram := s.metrics.NewDurationHistogram("GetAllUser_cache_histogram",
		"time for cache getAllUser",
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
	histogram := s.metrics.NewDurationHistogram("GetUserByEmail_endpoint_histogram",
		"time for cache getUserByEmail",
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
