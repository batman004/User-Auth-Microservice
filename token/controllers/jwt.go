package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)



var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT()( string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "test"
	claims["aud"] = "post.jwt.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute*1).Unix()

	tokenString, err := token.SignedString(MySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}