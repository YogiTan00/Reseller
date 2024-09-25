package delivery

import (
	"context"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/pkg/utils"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/internal/delivery/request"
	"github.com/YogiTan00/Reseller/services/transactions/internal/delivery/response"
	"github.com/google/uuid"
)

func (t TransactionHandler) GetListTransaction(ctx context.Context, req *transactionPb.GeneralFilter) (*transactionPb.TransactionList, error) {
	l := logger.Logger{
		EndPoint:    "/api/v1/transaction/list",
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()
	ctx = utils.SetCustomMetaDataTransactionId(ctx, l.TrxId)
	data := request.NewGeneralFilter(req)

	prd, err := t.transaction.GetListTransaction(ctx, data)
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := response.TransactionListResponse(prd)
	l.ResponseData = rsp
	return rsp, nil
}
