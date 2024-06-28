package routes

import (
	"fmt"
	"net/http"

	"github.com/dionyalcantara/go-rest-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Hello)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server")
		return
	}
}
