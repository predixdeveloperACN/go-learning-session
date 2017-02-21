package storage

import (
	"github.com/predixdeveloperACN/go-learning-session/model"
	"github.com/predixdeveloperACN/go-learning-session/postgres"
	"fmt"
)

func GetDogs(breed string) (returnDogs []model.Dog){
	err := postgres.Database.Select(&returnDogs, "SELECT * FROM animals.dogs WHERE breed = $1", breed)
	if err != nil {
		fmt.Println(err.Error())
	}
	return

}
