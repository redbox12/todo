package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/redbox12/todo-app/domain"
	"github.com/redbox12/todo-app/pkg/repository"
	"time"
)

const (
	salt      = "lakdjf213jklz3"
	signInKey = "fljkhds#" // ключ подписи для расшифровки
	tokenTTL  = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHas(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHas(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(), //генерация токена 12 часов
			IssuedAt:  time.Now().Unix(),               //время генерации токена
		},
		user.Id,
	})

	return token.SignedString([]byte(signInKey))
}

func (s *AuthService) generatePasswordHas(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}
		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, err
	}

	clailm, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Invalid token claims are not of type *tokenClaims")
	}

	return clailm.UserId, nil
}
