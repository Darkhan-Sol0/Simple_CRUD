package web

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (h *Handler) TimeoutAndSemoporeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), requestTimeout)
		defer cancel()
		if !h.checkSemaphore(c, ctx) {
			c.Abort()
			return
		}
		defer h.releaseSemaphore()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
