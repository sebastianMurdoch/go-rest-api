package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastianMurdoch/go-rest-api/src/api/domain/status"
	"net/http"
)

type StatusHandler struct {
	statusService status.StatusService
}

func NewStatusHandler(service status.StatusService) *StatusHandler{
	return &StatusHandler{statusService:service}
}

func (sh *StatusHandler) Get(c *gin.Context) {
	response, err := sh.statusService.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, response)
}