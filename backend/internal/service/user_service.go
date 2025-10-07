package service

import (
	"errors"

	"github.com/Gabrieldevreis/FornoGestor/internal/dto"
	"github.com/Gabrieldevreis/FornoGestor/internal/models"
	"github.com/Gabrieldevreis/FornoGestor/internal/repository"
	"github.com/Gabrieldevreis/FornoGestor/internal/utils"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil && existingUser.ID != 0 {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
		Active:   req.Active,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		Active: user.Active,
	}, nil
}

func (s *UserService) GetByID(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		Active: user.Active,
	}, nil
}

func (s *UserService) List() ([]dto.UserResponse, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = dto.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Role:   user.Role,
			Active: user.Active,
		}
	}

	return responses, nil
}

func (s *UserService) Update(id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" && req.Email != user.Email {
		existingUser, _ := s.repo.FindByEmail(req.Email)
		if existingUser != nil && existingUser.ID != id {
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.Role != "" {
		user.Role = req.Role
	}

	if req.Active != nil {
		user.Active = *req.Active
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		Active: user.Active,
	}, nil
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserService) ChangePassword(id uint, req dto.ChangePasswordRequest) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if !utils.CheckPassword(req.OldPassword, user.Password) {
		return errors.New("invalid old password")
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.repo.Update(user)
}
