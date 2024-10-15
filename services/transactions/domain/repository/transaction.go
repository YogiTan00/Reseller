package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
)

type TransactionRepositoryInterface interface {
	CreateTransaction(ctx context.Context, trans *entity.Transaction) error
	DeleteTransaction(ctx context.Context, transId string) error
	GetListTransaction(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.TransactionDto, error)
	GetDetailTransaction(ctx context.Context, filter *entity.GeneralFilter) (*entity.TransactionDto, error)
	UpdateTransaction(ctx context.Context, trans *entity.Transaction) error
	GetListTransactionCount(ctx context.Context, filter *entity.GeneralFilter) (int64, error)
}
