package models

import "time"

type Product struct {
	Id           string    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	TypeSize     string    `gorm:"column:type_size"`
	MarketType   string    `gorm:"column:market_type"`
	Image        string    `gorm:"column:image"`
	DefaultPrice int64     `gorm:"column:default_price"`
	Price        int64     `gorm:"column:price"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
