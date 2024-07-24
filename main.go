package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-rosinsky/go-cdn/handlers"
)

func main() {
	fmt.Println("Registering handles...")
	r := mux.NewRouter()
	r.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	r.HandleFunc("/{id}", handlers.RetrieveHandler).Methods("GET")

	fmt.Println("Serving on port 8080...")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
