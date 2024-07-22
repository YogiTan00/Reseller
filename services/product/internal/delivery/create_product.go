package delivery

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/google/uuid"
	"time"
)

func (prod *ProductHandler) CreateProduct(ctx context.Context, req *productPb.ProductRequest) (*productPb.ProductResponse, error) {
	l := logger.Logger{
		TimeStarted: time.Now(),
		EndPoint:    "/api/v1/product/create",
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()

	rsp := &productPb.ProductResponse{
		Id: req.GetId(),
	}
	return rsp, nil
}
