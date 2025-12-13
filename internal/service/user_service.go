package service

import (
	"crist-blog/internal/model"
	"crist-blog/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Login(username, password string) (*model.User, error) {
	user, err := s.userRepo.GetByName(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
