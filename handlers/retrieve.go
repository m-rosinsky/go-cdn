package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-rosinsky/go-cdn/storage"
)

func RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fileData, err := storage.LoadFile(id)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg") // Assuming jpeg.
	w.Write(fileData)
}
