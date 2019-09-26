package repositories

import (
	"github.com/smurdoch/my_simple_rest_api/src/api/domain/admission"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AdmissionRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdmissionRepository() admission.AdmissionRepository {
	db, err := gorm.Open("mysql", "spring_app:password@tcp(localhost:3306)/spring_database")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return &AdmissionRepositoryImpl{
		DB: db,
	}
}

func (admissionRepositoryImpl *AdmissionRepositoryImpl) GetAll() []admission.Admission {
	admissions := []admission.Admission{}

	admissionRepositoryImpl.DB.Find(&admissions)

	return admissions
}
