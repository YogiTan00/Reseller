package delivery

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery/request"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery/response"
	"github.com/google/uuid"
)

func (prod *ProductHandler) GetListProduct(ctx context.Context, req *productPb.GeneralFilter) (*productPb.ProductList, error) {
	l := logger.Logger{
		EndPoint:    "/api/v1/product/list",
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()

	data := request.NewGeneralFilter(req)

	prd, err := prod.product.GetListProduct(ctx, data)
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := response.ProductListResponse(prd)
	l.ResponseData = rsp
	return rsp, nil
}
