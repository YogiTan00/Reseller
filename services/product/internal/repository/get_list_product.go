package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/models"
)

func (p ProductRepository) GetListProduct(ctx context.Context, filter *entity.ProductFilter) ([]*entity.ProductDto, error) {
	var mdl []*models.Product
	txQuery := p.db.WithContext(ctx).Model(mdl)
	tx := p.txQueryProduct(txQuery, filter).Find(&mdl)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mapper.ToListEntityProduct(mdl), nil
}
