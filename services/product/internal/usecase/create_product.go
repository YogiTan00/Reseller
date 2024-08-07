package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

func (p ProductUsecase) CreateProduct(ctx context.Context, prd *entity.ProductDto) error {
	err := p.repoProduct.CreateProduct(ctx, prd.Create())
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
