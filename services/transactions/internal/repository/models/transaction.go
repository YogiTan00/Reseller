package models

import "time"

type Product struct {
	Id           string     `gorm:"column:id"`
	Name         string     `gorm:"column:name"`
	TypeSize     string     `gorm:"column:type_size"`
	Image        string     `gorm:"column:image"`
	DefaultPrice float64    `gorm:"column:default_price"`
	Price        float64    `gorm:"column:price"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at"`
}

func (Product) TableName() string {
	return "products"
}

type Transaction struct {
	Id          string    `gorm:"column:id"`
	IdName      string    `gorm:"column:id_name"`
	Return      *bool     `gorm:"column:return_trans"`
	SalesDate   time.Time `gorm:"column:sales_date"`
	Unit        int       `gorm:"column:unit"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	Product     Product   `gorm:"foreignKey:id_name;references:id"`
}

func (Transaction) TableName() string {
	return "transactions"
}
