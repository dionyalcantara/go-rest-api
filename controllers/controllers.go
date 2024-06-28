package controllers

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	fmt.Println("Hello :)")
}
