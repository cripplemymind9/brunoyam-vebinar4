package models

type Task struct {
	Id 				int 	`json:"id"`
	Title 			string 	`json:"title" validate:"required"`
	Description 	string 	`json:"descrription"`
	Status 			string 	`json:"status"`
}

type Response struct {
	Message 		string
}