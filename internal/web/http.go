package web

import (
	"MyProgy/internal/datasource"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Storage datasource.Storage
	sem     chan struct{}
}

func NewHandler(storage datasource.Storage) Handler {
	return Handler{
		Storage: storage,
		sem:     make(chan struct{}, 50),
	}
}

func (h *Handler) RegHandlers(r *gin.Engine) {
	r.GET("/", h.TimeoutAndSemoporeMiddleware(), h.MainHandler)
	r.GET("/hello", h.TimeoutAndSemoporeMiddleware(), h.HelloHandler)

	r.POST("/user", h.TimeoutAndSemoporeMiddleware(), h.CreateUserHandler)
	r.GET("/users", h.TimeoutAndSemoporeMiddleware(), h.GetUsersHandler)
	r.GET("/user/:id", h.TimeoutAndSemoporeMiddleware(), h.GetUserIdHandler)
	r.PATCH("/user/:id", h.TimeoutAndSemoporeMiddleware(), h.UpdateUserHandler)
	r.DELETE("user/:id", h.TimeoutAndSemoporeMiddleware(), h.DeleteUserHandler)
}

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
