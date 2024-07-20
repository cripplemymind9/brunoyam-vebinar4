package storage

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"
	"errors"
)

func (s *Storage) GetAllUsers() ([]models.User, error) {
	var users []models.User
	for _, user := range s.userDB.UserList {
		users = append(users, user)
	}
	return users, nil
}

func (s *Storage) InsertUser(user models.User) error {
	s.userDB.UserList[len(s.userDB.UserList)] = user
	return nil
}

func (s *Storage) GetUser(id int) (models.User, error) {
	return s.userDB.UserList[id-1], nil
}

func (s *Storage) UpdateUser(user models.User, id int) error {
	s.userDB.UserList[id-1] = user
	return nil
}

func (s *Storage) DeleteUser(id int) error {
	_, ok := s.userDB.UserList[id-1]
	if !ok {
		return errors.New("user not found")
	}
	s.userDB.UserList[id-1] = models.User{}
	return nil
}