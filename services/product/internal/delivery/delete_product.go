package delivery

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/google/uuid"
)

func (prod *ProductHandler) DeleteProduct(ctx context.Context, req *productPb.GeneralIdRequest) (*productPb.GeneralResponse, error) {
	l := logger.Logger{
		EndPoint:    fmt.Sprintf("/api/v1/product/delete/%s", req.GetId()),
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()

	err := prod.product.DeleteProduct(ctx, req.GetId())
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
