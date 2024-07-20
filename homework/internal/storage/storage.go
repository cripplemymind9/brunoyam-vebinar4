package storage

import "github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"

type TaskMap struct {
	TaskList map[int]models.Task
}

type UserMap struct {
	UserList map[int]models.User
}

type Storage struct {
	taskDB TaskMap
	userDB UserMap
}

func New() (*Storage, error) {
	taskDB := TaskMap{TaskList: make(map[int]models.Task)}
	userDB := UserMap{UserList: make(map[int]models.User)}
	return &Storage{
		taskDB: taskDB,
		userDB: userDB,
	}, nil
}