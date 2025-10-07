package dto

import "github.com/Gabrieldevreis/FornoGestor/internal/models"

type UserResponse struct {
	ID     uint            `json:"id"`
	Name   string          `json:"name"`
	Email  string          `json:"email"`
	Role   models.UserRole `json:"role"`
	Active bool            `json:"active"`
}

type CreateUserRequest struct {
	Name     string          `json:"name" binding:"required"`
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required,min=6"`
	Role     models.UserRole `json:"role" binding:"required,oneof=admin caixa garcom"`
	Active   bool            `json:"active"`
}

type UpdateUserRequest struct {
	Name   string          `json:"name"`
	Email  string          `json:"email" binding:"omitempty,email"`
	Role   models.UserRole `json:"role" binding:"omitempty,oneof=admin caixa garcom"`
	Active *bool           `json:"active"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
