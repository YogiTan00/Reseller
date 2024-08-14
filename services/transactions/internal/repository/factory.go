package repository

import (
	"github.com/YogiTan00/Reseller/services/transactions/domain/repository"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db gorm.DB
}

type TransactionRepositoryFactory struct {
	Db gorm.DB
}

func (prod *TransactionRepositoryFactory) Create() repository.TransactionRepositoryInterface {
	return &TransactionRepository{
		db: prod.Db,
	}
}
