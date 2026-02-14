package service

import (
	"wallet-app/model"
	"wallet-app/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
}

type IUserService interface {
	GetUserByID(userID uint) (*model.User, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByID(userID uint) (*model.User, error) {
	user, err := s.userRepository.FindUserByID(nil, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
