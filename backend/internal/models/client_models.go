package models

type Client struct {
	BaseModel
	Name    string  `gorm:"not null" json:"name"`
	Phone   string  `gorm:"index" json:"phone"`
	Email   string  `json:"email"`
	Address string  `json:"address"`
	CPF     string  `gorm:"index" json:"cpf"`
	Orders  []Order `gorm:"foreignKey:ClientID" json:"orders,omitempty"`
}
