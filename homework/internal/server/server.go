package server

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

type Storage interface {
	//Task
	GetAllTasks()([]models.Task, error)
	InsertTask(models.Task)error
	GetTask(int)(models.Task, error)
	UpdateTask(models.Task, int) error
	DeleteTask(int) error

	// User
	GetAllUsers()([]models.User, error)
	InsertUser(models.User)error
	GetUser(int)(models.User, error)
	UpdateUser(models.User, int) error
	DeleteUser(int) error

	//Auth
	Login(models.LoginUser) (int, error)
	Profile(models.Claims) (models.User, error)
}

type Server struct {
	addr 		string
	store 		Storage
	validate 	*validator.Validate
}

func NewServer(addr string, store Storage) *Server {
	return &Server{
		addr: 		addr,
		store: 		store,
		validate: 	validator.New(),
	}
}

func (s *Server) Run() error {
	router := gin.Default()

	router.POST("/users", s.InsertUserHandler)

	router.POST("/login", s.LoginHandler)
	router.GET("/profile", AuthMiddleware(), s.ProfileHandler)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(AuthMiddleware())
	{
		// Task routes routes
		protectedRoutes.GET("/tasks", s.GetAllTasksHanlder)
		protectedRoutes.GET("/tasks/:id", s.GetTaskHandler)
		protectedRoutes.POST("/tasks", s.InsertTaskHandler)
		protectedRoutes.PUT("/tasks/:id", s.UpdateTaskhandler)
		protectedRoutes.DELETE("/tasks/:id", s.DeleteTaskHanlder)
		//User protected routes
		protectedRoutes.GET("/users", s.GetAllUsersHanlder)
		protectedRoutes.GET("/users/:id", s.GetUserHandler)
		protectedRoutes.PUT("/users/:id", s.UpdateUserHanlder)
		protectedRoutes.DELETE("/users/:id", s.DeleteUserHandler)
	}

	return router.Run(s.addr)
}

func (s *Server) CreateToken(id int) (string, error) {
	user, err := s.store.GetUser(id)
	if err != nil  {
		return "", err
	}

	claims := &models.Claims{
		ID: 	user.ID,
		Email: 	user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret_key"))
}

func ValidateToken(ctx *gin.Context) (*models.Claims, error) {
	tokenString := ctx.GetHeader("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	
	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ValidateToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{Message: "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}