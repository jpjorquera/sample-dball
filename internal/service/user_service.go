package service

import (
	"dballz/internal/model"
	"dballz/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers() ([]model.User, error) {
	return s.repo.ListUsers()
}
