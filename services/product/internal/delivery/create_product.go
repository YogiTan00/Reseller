package delivery

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery/request"
	"github.com/google/uuid"
	"time"
)

func (prod *ProductHandler) CreateProduct(ctx context.Context, req *productPb.Product) (*productPb.Product, error) {
	l := logger.Logger{
		TimeStarted: time.Now(),
		EndPoint:    "/api/v1/product/create",
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()
	data := request.NewProductRequest(req)

	err := prod.product.CreateProduct(ctx, data)
	if err != nil {
		return nil, err
	}

	rsp := req
	l.ResponseData = rsp
	return rsp, nil
}
