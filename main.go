package main

import (
	"fmt"
	"github.com/dionyalcantara/go-rest-api/routes"
)

func main() {
	fmt.Println("Server is running on port 8080")
	routes.HandleRequest()
}
