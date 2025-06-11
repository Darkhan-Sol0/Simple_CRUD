package web

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	data any
	err  error
}

func sendError(c *gin.Context, status int, res Result) {
	c.JSON(status, gin.H{
		"status":  status,
		"error":   res.data,
		"details": res.err.Error(),
	})
}

func sendSuccess(c *gin.Context, status int, res Result) {
	c.JSON(status, gin.H{
		"status": status,
		"data":   res.data,
	})
}

func handleContextError(c *gin.Context, ctx context.Context) {
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		sendError(c, http.StatusGatewayTimeout, Result{data: "request timeout", err: nil})
	} else {
		sendError(c, http.StatusRequestTimeout, Result{data: "request cancelled", err: nil})
	}
}
