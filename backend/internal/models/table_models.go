package models

type TableStatus string

const (
	StatusAvailable TableStatus = "disponivel"
	StatusOccupied  TableStatus = "ocupada"
	StatusReserved  TableStatus = "reservada"
)

type Table struct {
	BaseModel
	Number   int         `gorm:"uniqueIndex;not null" json:"number"`
	Capacity int         `gorm:"not null" json:"capacity"`
	Status   TableStatus `gorm:"not null;default:'disponivel'" json:"status"`
	ClientID *uint       `json:"client_id"`
	Client   *Client     `gorm:"foreignKey:ClientID" json:"client,omitempty"`
	Orders   []Order     `gorm:"foreignKey:TableID" json:"orders,omitempty"`
}