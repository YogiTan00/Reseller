package usecase

import (
	"Reseller/pkg/logger"
	"Reseller/services/product/domain/repository"
	"Reseller/services/product/domain/usecase"
)

type ProductUsecase struct {
	l           logger.Logger
	repoProduct repository.ProductRepositoryInterface
}

type ProductUsecaseFactory struct {
	l           logger.Logger
	repoProduct repository.ProductRepositoryInterface
}

func (prod *ProductUsecaseFactory) Create() usecase.ProductUsecaseInterface {
	return &ProductUsecase{
		l:           prod.l,
		repoProduct: prod.repoProduct,
	}
}
