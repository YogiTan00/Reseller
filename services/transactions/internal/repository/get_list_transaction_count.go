package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/models"
)

func (t TransactionRepository) GetListTransactionCount(ctx context.Context, filter *entity.GeneralFilter) (int64, error) {
	var count int64
	filter.DisableOption()
	txQuery := t.db.WithContext(ctx).Table(models.Transaction{}.TableName())
	tx := t.txQueryTransaction(txQuery, filter).Count(&count)
	if tx.Error != nil {
		return int64(0), tx.Error
	}
	return count, nil
}
