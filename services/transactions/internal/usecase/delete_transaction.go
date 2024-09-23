package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/google/uuid"
)

func (t *TransactionUsecase) DeleteTransaction(ctx context.Context, transId string) error {
	err := uuid.Validate(transId)
	if err != nil {
		return err
	}

	filter := &entity.GeneralFilter{
		Id:        transId,
		IsDeleted: true,
	}
	trns, err := t.repoTransaction.GetDetailTransaction(ctx, filter)
	if err != nil {
		t.l.Error(err)
		return exceptions.ErrInternalServer
	}

	if trns == nil {
		return exceptions.ErrNotFound("transaction")
	}

	err = t.repoTransaction.DeleteTransaction(ctx, transId)
	if err != nil {
		t.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
