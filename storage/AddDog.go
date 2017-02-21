package storage

import (
	"github.com/predixdeveloperACN/go-learning-session/model"
	"github.com/predixdeveloperACN/go-learning-session/postgres"
)

func AddDog(dogObject model.Dog) (err error) {
	tx, err := postgres.Database.Begin()
	defer tx.Rollback()

	stmtDataRows, _ := tx.Prepare("INSERT INTO animals.dogs (breed,weight,height) VALUES ($1,$2,$3)")
	defer stmtDataRows.Close()
	stmtDataRows.Exec(dogObject.Breed, dogObject.Weight, dogObject.Height)
	tx.Commit()
	return
}
