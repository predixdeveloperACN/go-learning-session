package server

import (
	"github.com/predixdeveloperACN/go-learning-session/storage"
	"net/http"
)

func HandleDeleteDogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	breed := r.URL.Query().Get("breed")

	storage.DeleteDog(breed)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}