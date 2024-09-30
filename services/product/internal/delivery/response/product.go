package response

import (
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProductResponse(prd *entity.ProductDto) *productPb.Product {
	return &productPb.Product{
		Id:           prd.Id,
		Name:         prd.Name,
		TypeSize:     prd.TypeSize,
		Image:        prd.Image,
		DefaultPrice: uint64(prd.DefaultPrice),
		Price:        uint64(prd.Price),
		CreatedAt:    timestamppb.New(prd.CreatedAt),
		UpdatedAt:    timestamppb.New(prd.UpdatedAt),
	}
}

func ProductListResponse(prd []*entity.ProductDto) *productPb.ProductList {
	prdList := make([]*productPb.Product, 0)
	for _, v := range prd {
		prdList = append(prdList, ProductResponse(v))
	}

	return &productPb.ProductList{
		Data: prdList,
	}
}
