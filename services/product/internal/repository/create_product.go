package repository

import (
	"Reseller/services/product/domain/entity"
	"Reseller/services/product/internal/repository/mapper"
	"context"
)

func (p ProductRepository) CreateProduct(ctx context.Context, prd entity.Product) error {
	mdl := mapper.ToModelProduct(prd)
	tx := p.db.WithContext(ctx).Create(&mdl)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
