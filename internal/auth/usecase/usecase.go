package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"kienmatu/go-todos/internal/auth"
	"kienmatu/go-todos/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTL int64) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * time.Duration(tokenTTL),
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username, password string, limit int) (*models.User, error) {
	fmtusername := strings.ToLower(username)
	euser, _ := a.userRepo.GetUserByUsername(ctx, fmtusername)

	if euser != nil {
		return nil, auth.ErrUserExisted
	}
	user := &models.User{
		Id:       uuid.New().String(),
		Username: fmtusername,
		Password: password,
		Limit:    limit,
	}
	user.HashPassword()
	err := a.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return a.userRepo.GetUserByUsername(ctx, username)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, _ := a.userRepo.GetUserByUsername(ctx, username)
	if user != nil {
		return "", auth.ErrUserNotFound
	}

	if !user.ComparePassword(password) {
		return "", auth.ErrWrongPassword
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.expireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
