package delivery

import (
	"context"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery/response"
)

func (prod *ProductHandler) RpcGetProduct(ctx context.Context, req *productPb.GeneralIdRequest) (*productPb.Product, error) {
	prd, err := prod.product.GetDetailProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return response.ProductResponse(prd), nil
}
