package entity

import "time"

type Product struct {
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

type ProductDto struct {
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

func (data *Product) GetId() string {
	return data.id
}
func (data *Product) GetName() string {
	return data.name
}
func (data *Product) GetTypeSize() string {
	return data.typeSize
}
func (data *Product) GetMarketType() string {
	return data.marketType
}
func (data *Product) GetImage() string {
	return data.image
}
func (data *Product) GetDefaultPrice() int64 {
	return data.defaultPrice
}
func (data *Product) GetPrice() int64 {
	return data.price
}
func (data *Product) GetCreatedAt() time.Time {
	return data.createdAt
}
func (data *Product) GetUpdatedAt() time.Time {
	return data.updatedAt
}
