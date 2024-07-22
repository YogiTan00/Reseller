package repository

import (
	"github.com/YogiTan00/Reseller/services/product/domain/repository"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db gorm.DB
}

type ProductRepositoryFactory struct {
	Db gorm.DB
}

func (prod *ProductRepositoryFactory) Create() repository.ProductRepositoryInterface {
	return &ProductRepository{
		db: prod.Db,
	}
}
