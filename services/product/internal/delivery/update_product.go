package delivery

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery/request"
	"github.com/google/uuid"
)

func (prod *ProductHandler) UpdateProduct(ctx context.Context, req *productPb.Product) (*productPb.GeneralResponse, error) {
	l := logger.Logger{
		EndPoint:    "/api/v1/product/update",
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()
	data := request.UpdateProductRequest(req)

	err := prod.product.UpdateProduct(ctx, data)
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := &productPb.GeneralResponse{
		Data: &productPb.General{
			Message: "success",
		},
	}
	l.ResponseData = rsp
	return rsp, nil
}
