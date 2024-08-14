package models

import "time"

type Transaction struct {
	Id          string     `gorm:"column:id"`
	IdName      string     `gorm:"column:id_name"`
	Return      *bool      `gorm:"column:return_trans"`
	SalesDate   *time.Time `gorm:"column:sales_date"`
	Unit        int        `gorm:"column:unit"`
	Description string     `gorm:"column:description"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}
