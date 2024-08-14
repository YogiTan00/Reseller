package models

import "time"

type Product struct {
	Id           string     `gorm:"column:id"`
	Name         string     `gorm:"column:name"`
	TypeSize     string     `gorm:"column:type_size"`
	Image        string     `gorm:"column:image"`
	DefaultPrice int64      `gorm:"column:default_price"`
	Price        int64      `gorm:"column:price"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

func (Product) TableName() string {
	return "products"
}
