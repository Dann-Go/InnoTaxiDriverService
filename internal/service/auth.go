package service

import (
	"errors"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain"
	"github.com/Dann-Go/InnoTaxiDriverService/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	signingKey = "grk#iwoerjn%324Hsdj3skldflskdfn"
	tokenTTL   = 48 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	DriverId int `json:"driverId"`
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateDriver(driver *domain.Driver) (*domain.Driver, error) {
	driver.PasswordHash = hashPassword(driver.PasswordHash)

	return s.repo.CreateDriver(driver)
}

func (s *AuthService) GenerateToken(phone, password string) (string, error) {
	user, err := s.repo.GetDriverByPhone(phone)
	if err != nil {
		return "", err
	}
	if isValid := checkPasswordHash(password, user.PasswordHash); !isValid {
		err := errors.New("Wrong Password")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GetDriverByPhone(phone string) (*domain.Driver, error) {
	driver, err := s.repo.GetDriverByPhone(phone)
	if err != nil {
		return nil, err
	}

	return driver, err
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.DriverId, err
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
