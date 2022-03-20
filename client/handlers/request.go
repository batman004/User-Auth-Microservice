package handlers

import (
	"net/http"
	"log"
	"fmt"
	controller "client/controllers"

)

func HomePage(w http.ResponseWriter, _ *http.Request ){
	fmt.Fprintf(w, "Sensitive Database Info ")
}

func HandleRequests(){
	http.Handle("/", controller.IsAuthorized(HomePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}