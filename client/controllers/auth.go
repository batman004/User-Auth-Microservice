package controllers

import (
	"net/http"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
					return nil, fmt.Errorf(("Invalid signing method"))
				}
				aud := "post.jwt.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

				if !checkAudience {
					return nil, fmt.Errorf(("invalid aud"))
				}
				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return nil, fmt.Errorf(("invalid iss"))
				}

				return MySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid{
				endpoint(w,r)
			}
			
		}else{
			fmt.Fprintf(w, "No authorization token provided")
		}
	})
}