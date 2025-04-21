package response

import (
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func TransactionResponse(trns *entity.TransactionDto) *transactionPb.Transaction {
	var (
		returnItem *wrapperspb.BoolValue
	)
	if trns.ReturnItem != nil {
		returnItem = wrapperspb.Bool(*trns.ReturnItem)
	}
	return &transactionPb.Transaction{
		Id:          trns.Id,
		IdName:      trns.IdName,
		Name:        trns.Name,
		ReturnItem:  returnItem,
		SalesDate:   trns.SalesDate.Format(time.DateOnly),
		Unit:        uint32(trns.Unit),
		Description: trns.Description,
	}
}

func TransactionListResponse(prd []*entity.TransactionDto, count int64) *transactionPb.TransactionList {
	prdList := make([]*transactionPb.Transaction, 0)
	for _, v := range prd {
		prdList = append(prdList, TransactionResponse(v))
	}

	return &transactionPb.TransactionList{
		Data: prdList,
		Meta: &transactionPb.Metadata{
			Count: uint32(count),
		},
	}
}
