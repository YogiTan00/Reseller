package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/google/uuid"
)

func (p *ProductUsecase) DeleteProduct(ctx context.Context, productId string) error {
	err := uuid.Validate(productId)
	if err != nil {
		return err
	}

	filter := &entity.GeneralFilter{
		Id:        productId,
		IsDeleted: true,
	}
	prd, err := p.repoProduct.GetDetailProduct(ctx, filter)
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	if prd == nil {
		return exceptions.ErrNotFound("product")
	}
	err = p.repoProduct.DeleteProduct(ctx, productId)
	if err != nil {
		p.l.Error(err)
		return exceptions.ErrInternalServer
	}

	return nil
}
