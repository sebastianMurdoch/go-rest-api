package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
}

func NewStatusController() *StatusController {
	return &StatusController{}
}

func (controller StatusController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
