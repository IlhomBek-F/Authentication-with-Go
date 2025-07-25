package model

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Name  string   `json:"name"`
	Id    int      `json:"id"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}
