package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/models"
)

func (p ProductRepository) GetDetailProduct(ctx context.Context, productId string) (*entity.ProductDto, error) {
	var mdl *models.Product
	tx := p.db.WithContext(ctx).Where("id = ?", productId).First(&mdl)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mapper.ToEntityProduct(mdl), nil
}
