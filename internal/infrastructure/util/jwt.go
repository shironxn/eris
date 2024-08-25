package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shironxn/eris/internal/app/model"
)

type JWT struct {
	Access  string
	Refresh string
}

func NewJWT(jwt JWT) *JWT {
	return &JWT{
		Access:  jwt.Access,
		Refresh: jwt.Refresh,
	}
}

func (j *JWT) GenerateAccessToken(userID uint) (string, error) {
	exp := time.Now().Add(10 * time.Minute)
	claims := model.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Access))
}

func (j *JWT) GenerateRefreshToken(userID uint) (string, error) {

	fmt.Println(j.Refresh)
	exp := time.Now().Add(24 * time.Hour)
	claims := model.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Refresh))
}

func (j *JWT) ValidateToken(token string, secret string) (*model.Claims, error) {
	tokenString, err := jwt.ParseWithClaims(token, &model.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenString.Valid {
		return nil, errors.New("jwt token not valid")
	}

	claims, ok := tokenString.Claims.(*model.Claims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
