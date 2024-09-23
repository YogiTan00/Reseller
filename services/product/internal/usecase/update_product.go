package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/google/uuid"
)

func (p *ProductUsecase) UpdateProduct(ctx context.Context, prd *entity.ProductDto) error {
	err := uuid.Validate(prd.Id)
	if err != nil {
		return err
	}

	filter := &entity.GeneralFilter{
		Id: prd.Id,
	}
	prod, err := p.repoProduct.GetDetailProduct(ctx, filter)
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	if prod == nil {
		return exceptions.ErrNotFound("product")
	}

	err = p.repoProduct.UpdateProduct(ctx, prod.Update(prd))
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
