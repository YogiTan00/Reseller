package repository

import (
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"gorm.io/gorm"
)

func (p *ProductRepository) txQueryProduct(tx *gorm.DB, f *entity.GeneralFilter) *gorm.DB {
	if f != nil {
		if f.Id != "" {
			tx = tx.Where("id = ?", f.Id)
		}
		if f.IsDeleted {
			tx = tx.Where("deleted_at IS NULL")
		}
		if f.Q != "" {
			tx = tx.Where("name LIKE ?", "%"+f.Q+"%")
		}
		if f.Option != nil {
			var (
				sort, orderBy string
			)
			if f.Option.Sort != "" {
				sort = f.Option.Sort
			}
			if f.Option.OrderBy != "" {
				orderBy = f.Option.OrderBy
			}
			if sort != "" || orderBy != "" {
				tx = tx.Order(orderBy + " " + sort)
			}
			if f.Option.Limit > 0 {
				tx = tx.Limit(f.Option.Limit)
			}
			if f.Option.Offset > 0 {
				tx = tx.Offset(f.Option.Offset)
			}
		}
	}
	return tx

}
