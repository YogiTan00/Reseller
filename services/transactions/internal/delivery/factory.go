package delivery

import (
	"github.com/YogiTan00/Reseller/pkg/logger"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/domain/usecase"
	"google.golang.org/grpc"
)

type TransactionHandler struct {
	l           logger.Logger
	srv         *grpc.Server
	transaction usecase.TransactionUsecaseInterface
}

type TransactionHandlerFactory struct {
	L           logger.Logger
	Srv         *grpc.Server
	Transaction usecase.TransactionUsecaseInterface
}

func (prod *TransactionHandlerFactory) Create() *TransactionHandler {
	trans := &TransactionHandler{
		l:           prod.L,
		srv:         prod.Srv,
		transaction: prod.Transaction,
	}
	transactionPb.RegisterTransactionServiceServer(prod.Srv, trans)
	return trans
}
