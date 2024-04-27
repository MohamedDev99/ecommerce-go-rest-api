package token

import (
	"fmt"
	"time"

	"github.com/MohamedDev99/ecommerce-go-rest-api/pkg/configs"
	"github.com/golang-jwt/jwt/v5"
)

// logger
var log = configs.Logger

var secretKey = []byte("yR1geOK6BG7xFJyp7pw49zuz5UtvuzQGjGOvry4CeDA=")

// refresh token

// create token
func GenerateToken(email, firstName, lastName, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":     email,
			"firstName": firstName,
			"lastName":  lastName,
			"userId":    userId,
			"exp":       time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Error("Error generating token: " + err.Error())
		return "", err
	}

	return tokenString, nil
}

// verify token
func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Error("Error verifying token: " + err.Error())
		return err
	}

	if !token.Valid {
		log.Error("Invalid token")
		return fmt.Errorf("invalid token")
	}

	return nil
}
