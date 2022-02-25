package user

import (
	"errors"
	"log"
	"time"

	"cloud_storage/config"

	"github.com/dgrijalva/jwt-go"
)

func (user *User) EncodeJWT() (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":      "sgkandale-cloud-storage",
			"exp":      time.Now().Add(time.Minute * 120).Unix(),
			"iat":      time.Now().Unix(),
			"username": user.Username,
		},
	)

	tokenString, err := token.SignedString(
		[]byte(config.Config.Server.TokenSigningKey),
	)
	if err != nil {
		log.Println("user.EncodeJWT() : ", err)
		return "", errors.New("something went wrong")
	}

	return tokenString, nil
}
