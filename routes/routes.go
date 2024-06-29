package routes

import (
	"fmt"
	"net/http"

	"github.com/dionyalcantara/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Hello)
	router.HandleFunc("/characters", controllers.CreateCharacterHandler).Methods("POST")
	router.HandleFunc("/characters", controllers.GetCharactersHandler).Methods("GET")
	router.HandleFunc("/characters/{id}", controllers.GetCharacterByIdHandler).Methods("GET")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Error starting server")
		return
	}
}
