package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/mapper"
)

func (t TransactionRepository) CreateTransaction(ctx context.Context, trans *entity.Transaction) error {
	mdl := mapper.ToModelTransaction(trans)
	tx := t.db.WithContext(ctx).Create(&mdl)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
