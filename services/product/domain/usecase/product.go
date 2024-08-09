package usecase

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

type ProductUsecaseInterface interface {
	CreateProduct(ctx context.Context, prd *entity.ProductDto) error
	UpdateProduct(ctx context.Context, prd *entity.ProductDto) error
	DeleteProduct(ctx context.Context, productId string) error
	GetDetailProduct(ctx context.Context, productId string) (*entity.ProductDto, error)
	GetListProduct(ctx context.Context, filter *entity.GeneralFilter) ([]*entity.ProductDto, error)
}
