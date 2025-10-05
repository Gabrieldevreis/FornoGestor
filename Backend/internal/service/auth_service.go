package service

import (
	"errors"
	

	"github.com/Gabrieldevreis/FornoGestor/internal/config"
	"github.com/Gabrieldevreis/FornoGestor/internal/dto"
	"github.com/Gabrieldevreis/FornoGestor/internal/models"
	"github.com/Gabrieldevreis/FornoGestor/internal/repository"
	"github.com/Gabrieldevreis/FornoGestor/internal/utils"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repository.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.Active {
		return nil, errors.New("user is inactive")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := utils.GenerateToken(user, s.cfg.JWTSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.New().String()
	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Revoked:   false,
	}

	if err := s.userRepo.CreateRefreshToken(refreshTokenModel); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Role:   user.Role,
			Active: user.Active,
		},
	}, nil
}

func (s *AuthService) RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	refreshToken, err := s.userRepo.FindRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		return nil, errors.New("refresh token expired")
	}

	accessToken, err := utils.GenerateToken(&refreshToken.User, s.cfg.JWTSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.userRepo.RevokeRefreshToken(refreshToken)
}

func (s *AuthService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
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
