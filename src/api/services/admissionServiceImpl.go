package services

import (
	"github.com/smurdoch/my_simple_rest_api/src/api/domain/admission"
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
	return service.admissionRepository.GetAll()
}
