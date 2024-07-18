package repository

import (
	"context"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

type ProductRepositoryInterface interface {
	CreateProduct(ctx context.Context, prd entity.Product) error
	UpdateProduct(ctx context.Context, prd entity.Product) error
	DeleteProduct(ctx context.Context, productId string) error
	GetDetailProduct(ctx context.Context, productId string) (*entity.ProductDto, error)
	GetListProduct(ctx context.Context, filter *entity.ProductFilter) ([]*entity.ProductDto, error)
}
