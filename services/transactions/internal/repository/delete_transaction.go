package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/models"
	"time"
)

func (t *TransactionRepository) DeleteTransaction(ctx context.Context, transId string) error {
	tx := t.db.WithContext(ctx).Model(&models.Transaction{}).
		Where("id = ?", transId).
		Update("deleted_at", time.Now())
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
