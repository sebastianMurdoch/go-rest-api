package handlers

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"net/http"
)

type StatusHandler struct {
	Nr newrelic.Application `inject:"auto"`
}

func (sh *StatusHandler) Ping(c *gin.Context) {
	var txn newrelic.Transaction
	if sh.Nr != nil {
		txn = sh.Nr.StartTransaction("GET /ping", c.Writer, c.Request)
		defer txn.End()
	}
	c.String(http.StatusOK, "pong")
}