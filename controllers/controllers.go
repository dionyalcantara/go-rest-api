package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dionyalcantara/go-rest-api/database"
	"github.com/dionyalcantara/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	fmt.Println("Hello :)")
}

func CreateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var c models.Character
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.CreateCharacter(db, c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Character created")
}

func GetCharactersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	characters, err := models.GetCharacters(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(characters)
}

func GetCharacterByIdHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	character, err := models.GetCharacterById(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(character)
}
