package repository

import (
	"context"
	"errors"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/product/internal/repository/models"
	"gorm.io/gorm"
)

func (p *ProductRepository) GetDetailProduct(ctx context.Context, filter *entity.GeneralFilter) (*entity.ProductDto, error) {
	var mdl *models.Product
	txQuery := p.db.WithContext(ctx).Model(mdl)
	tx := p.txQueryProduct(txQuery, filter).First(&mdl)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return mapper.ToEntityProduct(mdl), nil
}
