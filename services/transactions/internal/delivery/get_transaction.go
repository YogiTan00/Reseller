package delivery

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/internal/delivery/response"
	"github.com/google/uuid"
)

func (t TransactionHandler) GetTransaction(ctx context.Context, req *transactionPb.GeneralIdRequest) (*transactionPb.Transaction, error) {
	l := logger.Logger{
		EndPoint:    fmt.Sprintf("/api/v1/transaction/view/%s", req.GetId()),
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()

	prd, err := t.transaction.GetDetailTransaction(ctx, req.GetId())
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := response.TransactionResponse(prd)
	l.ResponseData = rsp
	return rsp, nil
}
