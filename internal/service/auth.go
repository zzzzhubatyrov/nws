package service

import (
	"log"
	"nws/internal/model"
	"nws/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(repo repository.AuthRepositoryInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(user map[string]string) (*model.LoginResponse, error) {
	newUser := &model.User{
		Username: user["username"],
		Password: user["password"],
	}

	login, err := s.repo.Login(newUser)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(user["password"])); err != nil {
		log.Println("[service] invalid password:", err)
		return nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(login.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	token, err := claims.SignedString([]byte(viper.GetString("SECRET_KEY")))
	if err != nil {
		log.Println("[service] could not generate token:", err)
		return nil, err
	}

	return &model.LoginResponse{
		User:  login,
		Token: token,
	}, nil
}

func (s *AuthService) Register(user map[string]string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user["password"]), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Username: user["username"],
		Password: string(hashedPassword),
	}

	reg, err := s.repo.Register(newUser)
	if err != nil {
		return nil, err
	}

	return reg, nil
}
