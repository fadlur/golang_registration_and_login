package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWTKEY=[]byte("KEYnya Disini")

type JWTCLaim struct {
	Id int `json:"id"`
	Email string `json:"string"`
	jwt.StandardClaims
}

func GenerateJWT(id int, email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTCLaim{
		Id: id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(JWTKEY)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTCLaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWTKEY), nil
	},)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTCLaim)

	if !ok {
		err = errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}