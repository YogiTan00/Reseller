package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/mapper"
)

func (p *ProductRepository) CreateProduct(ctx context.Context, prd *entity.Product) error {
	mdl := mapper.ToModelProduct(prd)
	tx := p.db.WithContext(ctx).Create(&mdl)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
