package repository

import "github.com/cripplemymind9/brunoyam-vebinar4/classwork/internal/domain/models"

type Repository struct {
	db map[int]models.Task
}

func New() (*Repository, error) {
	db := make(map[int]models.Task)
	return &Repository{
		db: db,
	}, nil
}

func (repo *Repository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	for _, task := range repo.db {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repo *Repository) InsertTask(task models.Task) error {
	repo.db[len(repo.db)] = task
	return nil
}