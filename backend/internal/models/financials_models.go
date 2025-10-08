package models

import "time"

type EntryType string

const (
	TypeRevenue EntryType = "receita"
	TypeExpense EntryType = "despesa"
)

type FinancialCategory struct {
	BaseModel
	Name    string           `gorm:"not null;uniqueIndex" json:"name"`
	Type    EntryType        `gorm:"not null" json:"type"`
	Entries []FinancialEntry `gorm:"foreignKey:CategoryID" json:"entries,omitempty"`
}

type FinancialEntry struct {
	BaseModel
	Description string            `gorm:"not null" json:"description"`
	Amount      float64           `gorm:"not null" json:"amount"`
	Type        EntryType         `gorm:"not null" json:"type"`
	CategoryID  uint              `gorm:"not null" json:"category_id"`
	Category    FinancialCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Date        time.Time         `gorm:"not null" json:"date"`
	OrderID     *uint             `json:"order_id"`
	Order       *Order            `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	UserID      uint              `gorm:"not null" json:"user_id"`
	User        User              `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
