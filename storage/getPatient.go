package storage

import (
	"github.com/predixdeveloperACN/go-learning-session/model"
	"github.com/predixdeveloperACN/go-learning-session/postgres"
)

func GetPatients() (returnPatient []model.Patient, err error){
 	err = postgres.Database.Select(&returnPatient, "SELECT * FROM hospital.patients")
	return
}