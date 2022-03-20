package handlers

import (
	"fmt"
	"log"
	"net/http"
	controller "token/controllers"
)

func Index(w http.ResponseWriter, r *http.Request){
	validToken, err := controller.GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprintf(w, string(validToken))
}



func HandleRequests(){
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil ))
}