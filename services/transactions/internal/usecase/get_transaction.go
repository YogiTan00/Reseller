package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/google/uuid"
)

func (t TransactionUsecase) GetDetailTransaction(ctx context.Context, transId string) (*entity.TransactionDto, error) {
	err := uuid.Validate(transId)
	if err != nil {
		return nil, err
	}

	filter := &entity.GeneralFilter{
		Id:        transId,
		IsDeleted: true,
	}
	prd, err := t.repoTransaction.GetDetailTransaction(ctx, filter)
	if err != nil {
		t.l.Error(err)
		return nil, exceptions.ErrInternalServer
	}

	if prd == nil {
		return nil, exceptions.ErrNotFound("transaction")
	}

	return prd, nil
}
