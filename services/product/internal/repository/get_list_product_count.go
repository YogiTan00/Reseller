package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/models"
)

func (p *ProductRepository) GetListProductCount(ctx context.Context, filter *entity.GeneralFilter) (int64, error) {
	var count int64
	filter.DisableOption()
	txQuery := p.db.WithContext(ctx).Table(models.Product{}.TableName())
	tx := p.txQueryProduct(txQuery, filter).Count(&count)
	if tx.Error != nil {
		return int64(0), tx.Error
	}
	return count, nil
}
