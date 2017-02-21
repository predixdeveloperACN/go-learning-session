package server


import (
	"net/http"
	"github.com/predixdeveloperACN/go-learning-session/storage"
	"strconv"
)

func HandleUpdateDog(w http.ResponseWriter, r *http.Request) {
	weight,_ := strconv.Atoi(r.URL.Query().Get("weight"))
	breed := r.URL.Query().Get("breed")
	storage.UpdateDog(weight, breed)
}