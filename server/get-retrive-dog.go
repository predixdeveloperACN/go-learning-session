package server

import (
	"github.com/predixdeveloperACN/go-learning-session/storage"
	"net/http"
	"fmt"
	"encoding/json"
)

func HandleGetDogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	breed := r.URL.Query().Get("breed")

	dogs := storage.GetDogs(breed)
	fmt.Println(fmt.Sprintf("num dogs: %d", len(dogs)))
	outJSON, _ := json.Marshal(dogs)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(outJSON))
}
