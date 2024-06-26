package repository

import (
	"Reseller/services/product/internal/repository/models"
	"context"
	"time"
)

func (p ProductRepository) DeleteProduct(ctx context.Context, productId string) error {
	tx := p.db.WithContext(ctx).Model(&models.Product{}).
		Where("id = ?", productId).
		Update("deleted_at", time.Now())
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
