package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
)

func (t *TransactionUsecase) CreateTransaction(ctx context.Context, trans *entity.TransactionDto) error {
	err := trans.Validate()
	if err != nil {
		return err
	}

	err = t.repoTransaction.CreateTransaction(ctx, trans.New())
	if err != nil {
		t.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
