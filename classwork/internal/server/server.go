package server

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/classwork/internal/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Repository interface {
	GetAllTasks() ([]models.Task, error)
	InsertTask(models.Task) (error)
}

type Server struct {
	addr 		string
	db 			Repository
	validate 	*validator.Validate
}

func NewServer (addr string, db Repository) *Server {
	return &Server{
		addr: 		addr,
		db: 		db,
		validate: 	validator.New(),
	}
}

func (s *Server) Run() error {
	router := gin.Default()

	taksRoutes := router.Group("tasks")
	{
		taksRoutes.GET("/", s.GetTasksHandler)
		taksRoutes.POST("/", s.InsertTaskHandler)
	}

	return router.Run(s.addr)
}

func (s *Server) GetTasksHandler (ctx *gin.Context) {
	tasks, err := s.db.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (s *Server) InsertTaskHandler(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.validate.Struct(task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message":"validation failed"})
		return
	}

	if task.Status == "" {
		task.Status = "New"
	}

	if err := s.db.InsertTask(task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, "Task was saved!")
}