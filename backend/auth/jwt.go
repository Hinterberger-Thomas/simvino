package auth

import (
	"errors"
	"fmt"
	"simvino/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte(config.GetSecret())
)

type defaultClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken generates a jwt token and assign a email to it's claims and return it
func GenerateToken(email string) (string, error) {

	claims := defaultClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
			Issuer:    "simvino",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(SecretKey)
}

// ParseToken parses a jwt token and returns the email in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&defaultClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		},
	)
	if err != nil {
		return "", fmt.Errorf("val wrong")
	}

	claims, ok := token.Claims.(*defaultClaims)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}

	return claims.Email, nil
}
