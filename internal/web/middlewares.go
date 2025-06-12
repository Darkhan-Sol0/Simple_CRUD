package web

import (
	"MyProgy/pkg/jwt"
	"context"
	"fmt"
	"net/http"
	"strings"

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

func (h *Handler) ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			sendError(c, http.StatusUnauthorized, Result{data: "Unauth1", err: fmt.Errorf("Unauth1")})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			sendError(c, http.StatusUnauthorized, Result{data: "Unauth2", err: fmt.Errorf("Unauth2")})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := jwt.ValidateToken(tokenParts[1])
		if err != nil {
			sendError(c, http.StatusUnauthorized, Result{data: "Unauth3", err: fmt.Errorf("Unauth3")})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("claims", claims)

		c.Next()
	}
}
