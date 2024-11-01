package repository

import (
	"context"
	"errors"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/mapper"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository/models"
	"gorm.io/gorm"
)

func (t *TransactionRepository) GetDetailTransaction(ctx context.Context, filter *entity.GeneralFilter) (*entity.TransactionDto, error) {
	var mdl *models.Transaction
	txQuery := t.db.WithContext(ctx).Model(mdl)
	tx := t.txQueryTransaction(txQuery, filter).First(&mdl)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return mapper.ToEntityTransaction(mdl), nil
}
