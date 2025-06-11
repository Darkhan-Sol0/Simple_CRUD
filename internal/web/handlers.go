package web

import (
	"MyProgy/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) MainHandler(c *gin.Context) {
	ctx := c.Request.Context()
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		text := "Main Text"
		select {
		case resultChan <- Result{data: text, err: nil}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h Handler) HelloHandler(c *gin.Context) {
	ctx := c.Request.Context()
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		text := "Hello, World"
		select {
		case resultChan <- Result{data: text, err: nil}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h Handler) CreateUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "invalid request body", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		id, err := h.Storage.CreateUser(ctx, user)
		select {
		case resultChan <- Result{data: id, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "failed create user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusCreated, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h Handler) GetUsersHandler(c *gin.Context) {
	ctx := c.Request.Context()
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		users, err := h.Storage.GetUsers(ctx)
		select {
		case resultChan <- Result{data: users, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "failed get user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h Handler) GetUserIdHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "wrong id", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		user, err := h.Storage.GetUserId(ctx, int(id))
		select {
		case resultChan <- Result{data: user, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "failed get user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, res)
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
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
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		err := h.Storage.UpdateUser(ctx, int(id), user)
		select {
		case resultChan <- Result{data: nil, err: err}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "failed get user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, Result{data: "User updated", err: nil})
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}

func (h Handler) DeleteUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		sendError(c, http.StatusBadRequest, Result{data: "wrong id", err: err})
		return
	}
	resultChan := make(chan Result, 1)
	go func() {
		defer close(resultChan)
		res := h.Storage.DeleteUser(ctx, int(id))
		select {
		case resultChan <- Result{data: nil, err: res}:
		case <-ctx.Done():
			return
		}
	}()
	select {
	case res := <-resultChan:
		if res.err != nil {
			sendError(c, http.StatusInternalServerError, Result{data: "failed delete user", err: res.err})
			return
		}
		sendSuccess(c, http.StatusOK, Result{data: "User delete", err: nil})
	case <-ctx.Done():
		handleContextError(c, ctx)
		return
	}
}
