package product

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/pkg/utils"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"time"
)

func (p *ProductService) GetDetailProduct(trxId string, productId string) (*entity.ProductDto, error) {
	l := logger.Logger{
		EndPoint:    "GetDetailProduct",
		RequestData: productId,
		TrxId:       trxId,
	}
	defer l.CreateNewLog()
	ctx := context.WithValue(context.Background(), utils.HeaderTraceId, trxId)
	payloads := &productPb.GeneralIdRequest{Id: productId}
	prod, err := p.prdService.RpcGetProduct(ctx, payloads)
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	var deletedAt *time.Time
	if prod.GetDeletedAt() != nil {
		delAt := prod.GetDeletedAt().AsTime()
		deletedAt = &delAt
	}
	result := &entity.ProductDto{
		Id:           prod.GetId(),
		Name:         prod.GetName(),
		TypeSize:     prod.GetTypeSize(),
		Image:        prod.GetImage(),
		DefaultPrice: int64(prod.GetDefaultPrice()),
		Price:        int64(prod.GetPrice()),
		CreatedAt:    prod.GetCreatedAt().AsTime(),
		UpdatedAt:    prod.GetUpdatedAt().AsTime(),
		DeletedAt:    deletedAt,
	}

	l.ResponseData = result
	return result, nil
}
