package server

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) GetAllTasksHanlder(ctx *gin.Context) {
	tasks, err := s.store.GetAllTasks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (s *Server) InsertTaskHandler(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if task.Status == "" {
		task.Status = "New"
	}

	if err := s.validate.Struct(task); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.store.InsertTask(task); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "Task was saved"})
}

func (s *Server) GetTaskHandler(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	task, err := s.store.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (s *Server) UpdateTaskhandler(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	var task models.Task
	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.store.UpdateTask(task, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "Task was updated"})
}

func (s *Server) DeleteTaskHanlder(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	if err := s.store.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "Task was deleted"})
}