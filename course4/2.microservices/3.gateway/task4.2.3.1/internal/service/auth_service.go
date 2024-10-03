package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"metrics/internal/models"
)

type AuthRepository interface {
	Create(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (models.User, error)
}

type AuthServiceImpl struct {
	repo        AuthRepository
	token       *jwtauth.JWTAuth
	redisClient *redis.Client
}

func NewAuthServiceImpl(repo AuthRepository, token *jwtauth.JWTAuth, client *redis.Client) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo:        repo,
		token:       token,
		redisClient: client,
	}
}

func (a *AuthServiceImpl) CreateUser(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return a.repo.Create(context.Background(), user)
}

func (a *AuthServiceImpl) AuthUser(user models.User) (string, error) {
	userBD, err := a.repo.GetByEmail(context.Background(), user.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userBD.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("wrong password")
	}

	tokenStr, err := a.redisClient.Get(user.Email).Result()
	if err == nil {
		return tokenStr, nil
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	}
	_, token, _ := a.token.Encode(claims)
	err = a.redisClient.Set(user.Email, token, time.Hour*2).Err()
	if err != nil {
		log.Println(err)
	}
	return token, nil
}

func (a *AuthServiceImpl) VerifyToken(token string) (*models.User, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}
	jwtToken, err := a.token.Decode(token)
	if err != nil {
		return nil, err
	}

	res, err := a.redisClient.Get(jwtToken.PrivateClaims()["email"].(string)).Result()
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	if res != token {
		return nil, fmt.Errorf("invalid token")
	}
	user, err := a.repo.GetByEmail(context.Background(), jwtToken.PrivateClaims()["email"].(string))
	if err != nil {
		log.Println(err)
	}
	return &user, err
}
