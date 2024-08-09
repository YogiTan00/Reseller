package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

func (p ProductUsecase) GetListProduct(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.ProductDto, error) {
	filter.IsDeleted = true
	prd, err := p.repoProduct.GetListProduct(ctx, filter)
	if err != nil {
		p.l.Error(err)
		return nil, exceptions.ErrInternalServer
	}

	return prd, nil
}
