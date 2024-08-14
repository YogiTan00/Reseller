package response

import (
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TransactionResponse(trns *entity.TransactionDto) *transactionPb.Transaction {
	return &transactionPb.Transaction{
		Id:          trns.Id,
		IdName:      trns.IdName,
		Return:      wrapperspb.Bool(*trns.ReturnTrans),
		SalesDate:   timestamppb.New(*trns.SalesDate),
		Unit:        uint32(trns.Unit),
		Description: trns.Description,
	}
}

func TransactionListResponse(prd []*entity.TransactionDto) *transactionPb.TransactionList {
	prdList := make([]*transactionPb.Transaction, 0)
	for _, v := range prd {
		prdList = append(prdList, TransactionResponse(v))
	}

	return &transactionPb.TransactionList{
		Data: prdList,
	}
}
