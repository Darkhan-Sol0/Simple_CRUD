package web

import (
	"MyProgy/internal/datasource"

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
