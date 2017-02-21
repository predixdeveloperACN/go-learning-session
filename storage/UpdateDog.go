package storage

import "github.com/predixdeveloperACN/go-learning-session/postgres"

func UpdateDog(weight int, breed string) (err error) {
	tx, err := postgres.Database.Begin()
	defer tx.Rollback()

	stmtDataRows, _ := tx.Prepare("UPDATE animals.dogs SET weight = $1 WHERE breed = $2")
	defer stmtDataRows.Close()
	stmtDataRows.Exec(weight, breed)
	tx.Commit()
	return
}
