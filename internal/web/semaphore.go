package web

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	semTimeout     = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func (h *Handler) checkSemaphore(c *gin.Context, ctx context.Context) bool {
	select {
	case h.sem <- struct{}{}:
		return true
	case <-time.After(semTimeout):
		sendError(c, http.StatusTooManyRequests, Result{data: "service busy", err: nil})
		return false
	case <-ctx.Done():
		handleContextError(c, ctx)
		return false
	}
}

func (h *Handler) releaseSemaphore() {
	<-h.sem
}
