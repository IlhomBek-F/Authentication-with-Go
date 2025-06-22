package tokenutil

import (
	"auth/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(user model.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &model.JwtClaims{
		Name: user.Email,
		Id:   user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateRefreshToken() (refreshToken string, err error) {
	return "", nil
}

func IsAuthorized() (bool, error) {
	return false, nil
}
