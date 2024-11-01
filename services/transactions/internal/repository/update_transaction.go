package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/models"
)

func (t *TransactionRepository) UpdateTransaction(ctx context.Context, trans *entity.Transaction) error {
	mdl := mapper.ToModelTransaction(trans)
	tx := t.db.WithContext(ctx).Model(&models.Transaction{}).
		Where("id = ?", trans.GetId()).
		Save(&mdl)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
