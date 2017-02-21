package storage

import (
	"github.com/predixdeveloperACN/go-learning-session/postgres"
)

func DeleteDog(breed string) {
	tx, _ := postgres.Database.Begin()
	stmtDataRows, _ := tx.Prepare("DELETE FROM animals.dogs WHERE breed = $1")
	defer stmtDataRows.Close()
	stmtDataRows.Exec(breed)
	tx.Commit()
	return
}