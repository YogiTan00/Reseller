package usecase

import (
	"Reseller/services/product/domain/entity"
	"context"
)

func (p ProductUsecase) GetListProduct(ctx context.Context, filter *entity.ProductFilter) ([]*entity.ProductDto, error) {
	//TODO implement me
	panic("implement me")
}
