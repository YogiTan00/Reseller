package repository

import (
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"gorm.io/gorm"
)

func (p *TransactionRepository) txQueryTransaction(tx *gorm.DB, f *entity.GeneralFilter) *gorm.DB {
	if f != nil {
		if f.Id != "" {
			tx = tx.Where("id = ?", f.Id)
		}
		if f.IsDeleted {
			tx = tx.Where("deleted_at IS NULL")
		}
		if f.Q != "" {
			tx = tx.Where("name = ?", f.Q)
		}
	}
	return tx

}
