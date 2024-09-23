package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"github.com/google/uuid"
)

func (t *TransactionUsecase) GetDetailTransaction(ctx context.Context, transId string) (*entity.TransactionDto, error) {
	err := uuid.Validate(transId)
	if err != nil {
		return nil, err
	}

	filter := &entity.GeneralFilter{
		Id:        transId,
		IsDeleted: true,
	}
	trans, err := t.repoTransaction.GetDetailTransaction(ctx, filter)
	if err != nil {
		t.l.Error(err)
		return nil, exceptions.ErrInternalServer
	}

	if trans == nil {
		return nil, exceptions.ErrNotFound("transaction")
	}

	prod, err := t.serviceProduct.GetDetailProduct(t.l.TrxId, trans.IdName)
	if err != nil {
		t.l.Error(err)
	}
	if prod != nil {
		trans.Name = prod.Name
	} else {
		prod.Name = "(Unavailable)"
	}

	return trans, nil
}
