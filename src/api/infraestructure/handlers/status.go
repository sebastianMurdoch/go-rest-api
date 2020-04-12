package handlers

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"net/http"
)

type StatusHandler struct {
	nr newrelic.Application
}

func NewStatusHandler(nr newrelic.Application) *StatusHandler{
	return &StatusHandler{nr:nr}
}

func (sh *StatusHandler) Ping(c *gin.Context) {
	txn := sh.nr.StartTransaction("GET /ping", c.Writer, c.Request)
	defer txn.End()
	c.String(http.StatusOK, "pong")
}