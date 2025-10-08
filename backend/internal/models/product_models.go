package models

type ProductType string

const (
	TypePizza   ProductType = "pizza"
	TypeBebida  ProductType = "bebida"
	TypeLanche  ProductType = "lanche"
	TypeInsumo  ProductType = "insumo"
)

type ProductCategory struct {
	BaseModel
	Name        string    `gorm:"not null;uniqueIndex" json:"name"`
	Description string    `json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

type Product struct {
	BaseModel
	Name           string          `gorm:"not null" json:"name"`
	Description    string          `json:"description"`
	Price          float64         `gorm:"not null" json:"price"`
	Cost           float64         `json:"cost"`
	Type           ProductType     `gorm:"not null" json:"type"`
	CategoryID     *uint           `json:"category_id"`
	Category       ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Stock          int             `gorm:"default:0" json:"stock"`
	MinStock       int             `gorm:"default:0" json:"min_stock"`
	ImageURL       string          `json:"image_url"`
	Active         bool            `gorm:"default:true" json:"active"`
}