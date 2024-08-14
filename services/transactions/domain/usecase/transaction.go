package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
)

type TransactionUsecaseInterface interface {
	CreateTransaction(ctx context.Context, trans *entity.TransactionDto) error
	DeleteTransaction(ctx context.Context, transId string) error
	GetListTransaction(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.TransactionDto, error)
	GetDetailTransaction(ctx context.Context, transId string) (*entity.TransactionDto, error)
	UpdateTransaction(ctx context.Context, trans *entity.TransactionDto) error
}
