package storage

import (
	"github.com/cripplemymind9/brunoyam-vebinar4/homework/internal/domain/models"
	"errors"
)


func (s *Storage) Login(input models.LoginUser) (int, error) {
	for _, user := range s.userDB.UserList {
		if user.Email == input.Email && user.Password == input.Password {
			return user.ID, nil
		}
	}
	return -1, nil
}

func (s *Storage) Profile(claims models.Claims) (models.User, error) {
	for _, user := range s.userDB.UserList {
		if user.Email == claims.Email {
			return user, nil
		}
	}
	return models.User{}, errors.New("unauthorized")
}