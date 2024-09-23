package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/google/uuid"
)

func (t *TransactionUsecase) UpdateTransaction(ctx context.Context, trans *entity.TransactionDto) error {
	err := uuid.Validate(trans.Id)
	if err != nil {
		return err
	}

	filter := &entity.GeneralFilter{
		Id:        trans.Id,
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

	err = t.repoTransaction.UpdateTransaction(ctx, trns.Update(trans))
	if err != nil {
		t.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
