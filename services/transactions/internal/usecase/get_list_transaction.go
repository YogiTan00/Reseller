package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/utils"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
)

func (t *TransactionUsecase) GetListTransaction(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.TransactionDto, int64, error) {
	filter.IsDeleted = true
	prd, err := t.repoTransaction.GetListTransaction(ctx, filter)
	if err != nil {
		t.l.Error(err)
		return nil, int64(0), exceptions.ErrInternalServer
	}

	count, err := t.repoTransaction.GetListTransactionCount(ctx, filter)
	if err != nil {
		t.l.Error(err)
		return nil, int64(0), exceptions.ErrInternalServer
	}
	for _, value := range prd {
		prod, err := t.serviceProduct.GetDetailProduct(utils.GetTrxId(ctx), value.IdName)
		if err != nil {
			t.l.Error(err)
		}
		if prod != nil {
			value.Name = prod.Name
		} else {
			value.Name = "(Unavailable)"
		}
	}

	return prd, count, nil
}
