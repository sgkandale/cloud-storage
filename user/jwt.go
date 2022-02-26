package user

import (
	"errors"
	"fmt"
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
			"id":       user.Id,
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

type DecodedToken struct {
	Username string
	Id       int64
}

func DecodeToken(authToken string) (*DecodedToken, error) {

	if authToken == "" {
		return nil, errors.New("no auth token provided")
	}

	hmacSampleSecret := []byte(config.Config.Server.TokenSigningKey)

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &DecodedToken{
			Username: claims["username"].(string),
			Id:       int64(claims["id"].(float64)),
		}, nil
	}

	return nil, errors.New("invalid token")
}
