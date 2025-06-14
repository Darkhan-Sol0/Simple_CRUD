package web

import (
	models "MyProgy/internal/domain"
	"MyProgy/pkg/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) MainHandler(c *gin.Context) {
	text := "Main Text"
	sendSuccess(c, http.StatusOK, Result{data: text, err: nil})
}

func (h Handler) HelloHandler(c *gin.Context) {
	text := "Hello, World"
	sendSuccess(c, http.StatusOK, Result{data: text, err: nil})
}

func (h Handler) CreateUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	id, err := h.Storage.CreateUser(ctx, user)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "failed create user", err: err})
		return
	}
	sendSuccess(c, http.StatusCreated, Result{data: id, err: err})
}

func (h Handler) GetUsersHandler(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.Storage.GetUsers(ctx)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "failed get user", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: users, err: nil})

}

func (h Handler) GetUserIdHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "wrong id", err: err})
		return
	}
	user, err := h.Storage.GetUserById(ctx, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "failed get user", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: user, err: nil})
}

func (h Handler) UpdateUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "wrong id", err: err})
		return
	}
	var user models.User
	err = c.ShouldBindJSON(&user)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "failed update user", err: err})
		return
	}
	err = h.Storage.UpdateUser(ctx, int(id), user)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "failed get user", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: "User updated", err: nil})
}

func (h Handler) DeleteUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "wrong id", err: err})
		return
	}
	err = h.Storage.DeleteUser(ctx, int(id))
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "failed delete user", err: err})
		return
	}
	sendSuccess(c, http.StatusOK, Result{data: "User delete", err: nil})
}

func (h Handler) AuthUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	userOut, err := h.Storage.GetUserByName(ctx, user.Name, user.Password)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	token, err := jwt.GenerateToken(userOut.ID, userOut.Name, userOut.Role, userOut.Email)
	if err != nil {
		sendError(c, http.StatusNotFound, Result{data: "failed create user", err: err})
		return
	}
	sendSuccess(c, http.StatusCreated, Result{data: token, err: nil})
}
