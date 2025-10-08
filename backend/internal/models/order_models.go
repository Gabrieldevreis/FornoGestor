package models

type OrderStatus string

const (
	OrderPending   OrderStatus = "pendente"
	OrderPreparing OrderStatus = "preparando"
	OrderReady     OrderStatus = "pronto"
	OrderDelivered OrderStatus = "entregue"
	OrderCancelled OrderStatus = "cancelado"
	OrderClosed    OrderStatus = "fechado"
)

type Order struct {
	BaseModel
	TableID    uint        `gorm:"not null" json:"table_id"`
	Table      Table       `gorm:"foreignKey:TableID" json:"table,omitempty"`
	ClientID   *uint       `json:"client_id"`
	Client     *Client     `gorm:"foreignKey:ClientID" json:"client,omitempty"`
	UserID     uint        `gorm:"not null" json:"user_id"`
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Status     OrderStatus `gorm:"not null;default:'pendente'" json:"status"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	Total      float64     `json:"total"`
	Discount   float64     `gorm:"default:0" json:"discount"`
	FinalTotal float64     `json:"final_total"`
	Notes      string      `json:"notes"`
}

type OrderItem struct {
	BaseModel
	OrderID   uint    `gorm:"not null" json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	UnitPrice float64 `gorm:"not null" json:"unit_price"`
	Total     float64 `json:"total"`
	Notes     string  `json:"notes"`
}
