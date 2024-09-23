package entity

import "time"

type ProductDto struct {
	Id           string
	Name         string
	TypeSize     string
	Image        string
	DefaultPrice int64
	Price        int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
