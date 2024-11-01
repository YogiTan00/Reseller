package request

import (
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/domain/entity"
	"time"
)

func NewTransactionRequest(req *transactionPb.Transaction) *entity.TransactionDto {
	var (
		returnTrans *bool
		salesDate   *time.Time
	)
	if req.GetReturnItem() != nil {
		trans := req.GetReturnItem().GetValue()
		returnTrans = &trans
	}
	if req.GetSalesDate() != nil {
		sales := req.GetSalesDate().AsTime()
		salesDate = &sales
	}

	return &entity.TransactionDto{
		IdName:      req.GetIdName(),
		ReturnItem:  returnTrans,
		SalesDate:   salesDate,
		Unit:        int(req.GetUnit()),
		Description: req.GetDescription(),
	}
}

func UpdateTransactionRequest(req *transactionPb.Transaction) *entity.TransactionDto {
	var (
		returnTrans *bool
		salesDate   *time.Time
	)
	if req.GetReturnItem() != nil {
		trans := req.GetReturnItem().GetValue()
		returnTrans = &trans
	}
	if req.GetSalesDate() != nil {
		sales := req.GetSalesDate().AsTime()
		salesDate = &sales
	}

	return &entity.TransactionDto{
		Id:          req.GetId(),
		IdName:      req.GetIdName(),
		ReturnItem:  returnTrans,
		SalesDate:   salesDate,
		Unit:        int(req.GetUnit()),
		Description: req.GetDescription(),
	}
}

func NewGeneralFilter(req *transactionPb.GeneralFilter) *entity.GeneralFilter {
	return &entity.GeneralFilter{
		Q: req.GetQ(),
		Option: &entity.GeneralFilterOption{
			OrderBy: req.GetOrderBy(),
			Limit:   int(req.GetLimit()),
			Offset:  int(req.GetOffset()),
			Sort:    req.GetSort(),
		},
	}

}
