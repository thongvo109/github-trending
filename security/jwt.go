package security

import (
	"github-trending/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const KEY = "asdasdsahdkjasdsa"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
