package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/models"
)

func (t TransactionRepository) GetListTransaction(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.TransactionDto, error) {
	var mdl []*models.Transaction
	txQuery := t.db.WithContext(ctx).Model(mdl)
	tx := t.txQueryTransaction(txQuery, filter).Find(&mdl)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mapper.ToListEntityProduct(mdl), nil
}
