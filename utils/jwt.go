package utils

import (
	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("secret")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func GenerateToken(userId int) (string, error) {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 0,
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return token, claims, err
}
