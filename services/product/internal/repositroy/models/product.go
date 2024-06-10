package models

import "time"

type Product struct {
	Id           string `gorm:"column:id"`
	Name         string
	TypeSize     string
	MarketType   string
	Image        string
	DefaultPrice int64
	Price        int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
