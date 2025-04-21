package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/utils"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"strings"
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

	result := make([]*entity.TransactionDto, 0)
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
		if filter.Q != "" {
			if strings.Contains(strings.ToLower(value.Name), strings.ToLower(filter.Q)) {
				result = append(result, value)
			}
			continue
		} else {
			result = append(result, value)
		}
	}

	return result, count, nil
}
