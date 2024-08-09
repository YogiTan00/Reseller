package delivery

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery/response"
	"github.com/google/uuid"
)

func (prod *ProductHandler) GetProduct(ctx context.Context, req *productPb.GeneralIdRequest) (*productPb.Product, error) {
	l := logger.Logger{
		EndPoint:    fmt.Sprintf("/api/v1/product/view/%s", req.GetId()),
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()

	prd, err := prod.product.GetDetailProduct(ctx, req.GetId())
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := response.ProductResponse(prd)
	l.ResponseData = rsp
	return rsp, nil
}
