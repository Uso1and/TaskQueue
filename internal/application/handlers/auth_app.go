package handlers

import (
	"context"
	"errors"
	"time"

	"taskqueue/internal/domain/entities"
	"taskqueue/internal/domain/repositories"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthApp struct {
	userRepo  repositories.UserRepository
	jwtSecret []byte
}

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthApp(userRepo repositories.UserRepository, jwtSecret string) *AuthApp {
	return &AuthApp{
		userRepo:  userRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (a *AuthApp) Login(ctx context.Context, email, password string) (*entities.User, string, error) {
	user, err := a.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", errors.New("пользователь не найден")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("неверный пароль")
	}

	token, err := a.generateToken(user)
	if err != nil {
		return nil, "", errors.New("ошибка создания токена")
	}

	return user, token, nil
}

func (a *AuthApp) generateToken(user *entities.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.jwtSecret)
}

func (a *AuthApp) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return a.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("недействительный токен")
}
