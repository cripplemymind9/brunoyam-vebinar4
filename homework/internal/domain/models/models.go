package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	ID 			int 	`json:"id"`
	Email 		string 	`json:"email" validate:"email"`
	jwt.StandardClaims
}

type Response struct {
	Message string `json:"message"`
}