package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

func (p *ProductUsecase) GetListProduct(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.ProductDto, int64, error) {
	filter.IsDeleted = true
	prd, err := p.repoProduct.GetListProduct(ctx, filter)
	if err != nil {
		p.l.Error(err)
		return nil, int64(0), exceptions.ErrInternalServer
	}

	count, err := p.repoProduct.GetListProductCount(ctx, filter)
	if err != nil {
		p.l.Error(err)
		return nil, int64(0), exceptions.ErrInternalServer
	}

	return prd, count, nil
}
