package storage

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"
	"errors"
)

func (s *Storage) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	for _, task := range s.taskDB.TaskList {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Storage) InsertTask(task models.Task) error {
	s.taskDB.TaskList[len(s.taskDB.TaskList)] = task
	return nil
}

func (s *Storage) GetTask(id int) (models.Task, error) {
	return s.taskDB.TaskList[id-1], nil
}

func (s *Storage) UpdateTask(task models.Task, id int) error {
	s.taskDB.TaskList[id-1] = task
	return nil
}

func (s *Storage) DeleteTask(id int) error {
	_, ok := s.taskDB.TaskList[id-1]
	if !ok {
		return errors.New("task not found")
	}
	s.taskDB.TaskList[id-1] = models.Task{}
	return nil
}