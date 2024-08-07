package request

import (
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

func NewProductRequest(req *productPb.Product) *entity.ProductDto {
	return &entity.ProductDto{
		Name:         req.GetName(),
		TypeSize:     req.GetTypeSize(),
		MarketType:   req.GetMarketType(),
		Image:        req.GetImage(),
		DefaultPrice: int64(req.GetDefaultPrice()),
		Price:        int64(req.GetPrice()),
	}
}
