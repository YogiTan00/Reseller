package delivery

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/exceptions"
	"github.com/YogiTan00/Reseller/pkg/logger"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/google/uuid"
)

func (t TransactionHandler) DeleteTransaction(ctx context.Context, req *transactionPb.GeneralIdRequest) (*transactionPb.GeneralResponse, error) {
	l := logger.Logger{
		EndPoint:    fmt.Sprintf("/api/v1/transaction/delete/%s", req.GetId()),
		RequestData: req,
		TrxId:       uuid.New().String(),
	}
	defer l.CreateNewLog()

	err := t.transaction.DeleteTransaction(ctx, req.GetId())
	if err != nil {
		l.StatusCode = exceptions.MapToGrpcStatusCode(err)
		l.ResponseData = err.Error()
		return nil, exceptions.MapToGrpcStatusError(err)
	}

	rsp := &transactionPb.GeneralResponse{
		Data: &transactionPb.General{
			Message: "success",
		},
	}
	l.ResponseData = rsp
	return rsp, nil
}
