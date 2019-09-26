package services

import (
	"fmt"

	"github.com/smurdoch/go-rest-api/src/api/domain/admission"
)

type AdmissionServiceImpl struct {
	admissionRepository admission.AdmissionRepository
}

func NewAdmissionService(admissionRepository admission.AdmissionRepository) admission.AdmissionService {
	return &AdmissionServiceImpl{
		admissionRepository: admissionRepository,
	}
}

func (service *AdmissionServiceImpl) GetAll() []admission.Admission {
	fmt.Println("Service")
	return service.admissionRepository.GetAll()
}
