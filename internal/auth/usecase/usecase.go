package usecase

import (
	"context"
	"fmt"
	"kienmatu/go-todos/models"
	"kienmatu/go-todos/modules/auth"
	"strings"
	"time"

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
	tokenTTL time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username, password string, limit int) error {
	fmtusername := strings.ToLower(username)
	euser, _ := a.userRepo.GetUserByUsername(ctx, fmtusername)

	if euser != nil {
		return auth.ErrUserExisted
	}
	user := &models.User{
		Id: uuid.New(),
		Username: fmtusername,
		Password: password,
		Limit: limit,
	}
	user.HashPassword()
	return a.userRepo.CreateUser(ctx, user)
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
