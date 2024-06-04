package entity

import "time"

type Product struct {
	Id           string
	Name         string
	TypeSize     string
	MarketType   string
	Image        string
	DefaultPrice int64
	Price        int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductDto struct {
	id           string
	name         string
	typeSize     string
	marketType   string
	image        string
	defaultPrice int64
	price        int64
	createdAt    time.Time
	updatedAt    time.Time
}
