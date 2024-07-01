package routes

import (
	"fmt"
	"net/http"

	"github.com/dionyalcantara/go-rest-api/controllers"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func HandleRequest() {
	router := mux.NewRouter()

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	router.Use(handlers.CORS(headersOk, originsOk, methodsOk))

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
