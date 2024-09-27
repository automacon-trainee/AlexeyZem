package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"project/internal/auth/models"
	"project/internal/myerror"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
}

type AuthServiceImpl struct {
	repo        UserRepository
	token       *jwtauth.JWTAuth
	redisClient *redis.Client
}

func NewAuthServiceImpl(repo UserRepository, token *jwtauth.JWTAuth, client *redis.Client) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo:        repo,
		token:       token,
		redisClient: client,
	}
}

func (a *AuthServiceImpl) CreateUser(user models.User) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hash)
	err = a.repo.Create(context.Background(), &user)

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (a *AuthServiceImpl) AuthUser(user models.User) (string, error) {
	userBD, err := a.repo.GetByEmail(context.Background(), user.Email)
	if err != nil {
		return "", fmt.Errorf("AuthService :%w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(userBD.Password), []byte(user.Password))
	if err != nil {
		return "", myerror.ErrWrongPassword
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
		return nil, myerror.ErrWrongToken
	}
	jwtToken, err := a.token.Decode(token)
	if err != nil {
		return nil, err
	}

	res, err := a.redisClient.Get(jwtToken.PrivateClaims()["email"].(string)).Result()
	if err != nil {
		return nil, myerror.ErrWrongToken
	}

	if res != token {
		return nil, myerror.ErrWrongToken
	}
	user, err := a.repo.GetByEmail(context.Background(), jwtToken.PrivateClaims()["email"].(string))
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("VerifyToken :%w", err)
	}

	return &user, nil
}
