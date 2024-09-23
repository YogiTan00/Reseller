package usecase

import (
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
)

type ProductServiceInterface interface {
	GetDetailProduct(trxId string, productId string) (*entity.ProductDto, error)
}
