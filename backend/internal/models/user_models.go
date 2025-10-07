package models

import "time"

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleCaixa  UserRole = "caixa"
	RoleGarcom UserRole = "garcom"
)

type User struct {
	Name     string   `gorm:"not null" json:"name"`
	Email    string   `gorm:"uniqueIndex;not null" json:"email"`
	Password string   `gorm:"not null" json:"-"`
	Role     UserRole `gorm:"not null;default:'garcom'" json:"role"`
	Active   bool     `gorm:"default:true" json:"active"`
	BaseModel
}

type RefreshToken struct {
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	BaseModel
}
