package repository

import (
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"gorm.io/gorm"
)

func (p *ProductRepository) txQueryProduct(tx *gorm.DB, f *entity.ProductFilter) *gorm.DB {
	if f != nil {
		if f.Id != "" {
			tx = tx.Where("id = ?", f.Id)
		}
		if !f.IsDeleted {
			tx = tx.Where("deleted_at IS NOT NULL")
		}
	}
	return tx

}
