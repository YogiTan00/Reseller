package request

import (
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
)

func NewProductRequest(req *productPb.Product) *entity.ProductDto {
	return &entity.ProductDto{
		Name:         req.GetName(),
		TypeSize:     req.GetTypeSize(),
		Image:        req.GetImage(),
		DefaultPrice: int64(req.GetDefaultPrice()),
		Price:        int64(req.GetPrice()),
	}
}

func UpdateProductRequest(req *productPb.Product) *entity.ProductDto {
	return &entity.ProductDto{
		Id:           req.GetId(),
		Name:         req.GetName(),
		TypeSize:     req.GetTypeSize(),
		Image:        req.GetImage(),
		DefaultPrice: int64(req.GetDefaultPrice()),
		Price:        int64(req.GetPrice()),
	}
}

func NewGeneralFilter(req *productPb.GeneralFilter) *entity.GeneralFilter {
	return &entity.GeneralFilter{
		Q: req.GetQ(),
		Option: entity.GeneralFilterOption{
			OrderBy: req.GetOrderBy(),
			Limit:   int(req.GetLimit()),
			Offset:  int(req.GetOffset()),
		},
	}

}
