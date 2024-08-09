package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

func (p ProductUsecase) CreateProduct(ctx context.Context, prd *entity.ProductDto) error {
	err := prd.Validate()
	if err != nil {
		return err
	}

	err = p.repoProduct.CreateProduct(ctx, prd.New())
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
