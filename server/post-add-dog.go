package server

import (
	"net/http"
	"github.com/predixdeveloperACN/go-learning-session/storage"
	"github.com/predixdeveloperACN/go-learning-session/model"
	"strconv"
)

func HandleAddDog(w http.ResponseWriter, r *http.Request) {
	breed := r.URL.Query().Get("breed")
	weight,_ := strconv.Atoi(r.URL.Query().Get("weight"))
	height,_ := strconv.Atoi(r.URL.Query().Get("height"))

	dog := model.Dog{
		Breed: breed,
		Weight: weight,
		Height: height,
	}
	storage.AddDog(dog)
}