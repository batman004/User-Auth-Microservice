package main

import (
	"fmt"
	handler "client/handlers"
)

func main() {
	fmt.Println("server is running 🚄")
	handler.HandleRequests()
}

