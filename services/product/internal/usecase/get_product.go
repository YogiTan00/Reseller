package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"github.com/google/uuid"
)

func (p *ProductUsecase) GetDetailProduct(ctx context.Context, productId string) (*entity.ProductDto, error) {
	err := uuid.Validate(productId)
	if err != nil {
		return nil, err
	}

	filter := &entity.GeneralFilter{
		Id:        productId,
		IsDeleted: true,
	}
	prd, err := p.repoProduct.GetDetailProduct(ctx, filter)
	if err != nil {
		p.l.Error(err)
		return nil, exceptions.ErrInternalServer
	}

	if prd == nil {
		return nil, exceptions.ErrNotFound("product")
	}

	return prd, nil
}
