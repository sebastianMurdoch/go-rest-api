package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/smurdoch/go-rest-api/src/api/domain/admission"
)

type AdmissionController struct {
	admissionService admission.AdmissionService
}

func NewAdmissionController(service admission.AdmissionService) *AdmissionController {
	return &AdmissionController{
		admissionService: service,
	}
}

func (controller *AdmissionController) GetAll(c *gin.Context) {
	c.JSON(200, controller.admissionService.GetAll())
}
