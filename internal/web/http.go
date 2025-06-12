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
	r.GET("/auth", h.TimeoutAndSemoporeMiddleware(), h.AuthUserHandler)
	r.POST("/user", h.TimeoutAndSemoporeMiddleware(), h.CreateUserHandler)

	authGroup := r.Group("/", h.ValidateTokenMiddleware())
	{
		// authGroup.POST("/user", h.TimeoutAndSemoporeMiddleware(), h.CreateUserHandler)
		authGroup.GET("/users", h.TimeoutAndSemoporeMiddleware(), h.GetUsersHandler)
		authGroup.GET("/user/:id", h.TimeoutAndSemoporeMiddleware(), h.GetUserIdHandler)
		authGroup.PATCH("/user/:id", h.TimeoutAndSemoporeMiddleware(), h.UpdateUserHandler)
		authGroup.DELETE("user/:id", h.TimeoutAndSemoporeMiddleware(), h.DeleteUserHandler)
	}

}
