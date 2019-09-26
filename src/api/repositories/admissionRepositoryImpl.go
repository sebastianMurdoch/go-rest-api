package repositories

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/smurdoch/go-rest-api/src/api/domain/admission"
	"github.com/smurdoch/go-rest-api/src/api/persistence"

	_ "github.com/mattn/go-sqlite3"
)

type AdmissionRepositoryImpl struct {
	db *sqlx.DB
}

func NewAdmissionRepository() admission.AdmissionRepository {

	db, _ := persistence.DB()

	return &AdmissionRepositoryImpl{
		db: db,
	}
}

func (admissionRepositoryImpl *AdmissionRepositoryImpl) GetAll() []admission.Admission {
	admissions := []admission.Admission{}

	rows, err := admissionRepositoryImpl.db.Queryx("SELECT * FROM admissions")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var row admission.Admission
		err = rows.StructScan(&row)
		if err != nil {
			log.Fatal(err)
		}
		admissions = append(admissions, row)
		fmt.Println("Select item!")
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return admissions
}
