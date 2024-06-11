package repositroy

import (
	"Reseller/services/product/domain/entity"
	"Reseller/services/product/internal/repositroy/mapper"
	"Reseller/services/product/internal/repositroy/models"
	"context"
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
