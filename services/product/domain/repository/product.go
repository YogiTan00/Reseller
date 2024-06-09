package repository

import (
	"Reseller/services/product/domain/entity"
	"context"
)

type ProductRepositoryInterface interface {
	CreateProduct(ctx context.Context, prd entity.Product) error
	UpdateProduct(ctx context.Context, prd entity.Product) error
	DeleteProduct(ctx context.Context, productId string) error
	GetProduct(ctx context.Context, productId string) error
}
