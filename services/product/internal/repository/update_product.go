package repository

import (
	"Reseller/services/product/domain/entity"
	"Reseller/services/product/internal/repositroy/mapper"
	"Reseller/services/product/internal/repositroy/models"
	"context"
)

func (p ProductRepository) UpdateProduct(ctx context.Context, prd entity.Product) error {
	mdl := mapper.ToModelProduct(prd)
	tx := p.db.WithContext(ctx).Model(&models.Product{}).
		Where("id = ?", prd.GetId()).
		Save(&mdl)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
