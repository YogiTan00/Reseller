package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/models"
)

func (p *ProductRepository) UpdateProduct(ctx context.Context, prd entity.Product) error {
	mdl := mapper.ToModelProduct(prd)
	tx := p.db.WithContext(ctx).Model(&models.Product{}).
		Where("id = ?", prd.GetId()).
		Save(&mdl)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
