package main

import (
	"fmt"
	"github.com/dionyalcantara/go-rest-api/routes"
	"github.com/dionyalcantara/go-rest-api/database"
)

func main() {
	fmt.Println("Server is running on port 8080")
	_, err := database.InitDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	routes.HandleRequest()
}
