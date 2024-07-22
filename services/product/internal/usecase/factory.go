package usecase

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/services/product/domain/repository"
	"github.com/YogiTan00/Reseller/services/product/domain/usecase"
)

type ProductUsecase struct {
	l           logger.Logger
	repoProduct repository.ProductRepositoryInterface
}

type ProductUsecaseFactory struct {
	L           logger.Logger
	RepoProduct repository.ProductRepositoryInterface
}

func (prod *ProductUsecaseFactory) Create() usecase.ProductUsecaseInterface {
	return &ProductUsecase{
		l:           prod.L,
		repoProduct: prod.RepoProduct,
	}
}
