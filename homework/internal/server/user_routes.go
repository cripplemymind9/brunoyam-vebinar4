package server

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) GetAllUsersHanlder(ctx *gin.Context) {
	users, err := s.store.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (s *Server) InsertUserHandler(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.store.InsertUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "User was saved"})
}

func (s *Server) GetUserHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	user, err := s.store.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (s *Server) UpdateUserHanlder(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	var user models.User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.store.UpdateUser(user, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "User was updated"})
}

func (s *Server) DeleteUserHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.store.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "User was deleted"})
}