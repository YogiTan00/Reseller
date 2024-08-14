package cmd

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/logger"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/internal/delivery"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository"
	"github.com/YogiTan00/Reseller/services/transactions/internal/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

type TransactionInit struct {
	Db  *gorm.DB
	Srv *grpc.Server
	L   logger.Logger
}

func (prod *TransactionInit) Create(ctx context.Context, muxHttp *runtime.ServeMux, port string, opts []grpc.DialOption) {
	prod.L.EndPoint = "Transaction"
	repoTransaction := &repository.TransactionRepositoryFactory{
		Db: *prod.Db,
	}
	ucTransaction := &usecase.TransactionUsecaseFactory{
		L:               prod.L,
		RepoTransaction: repoTransaction.Create(),
	}
	srvTransaction := &delivery.TransactionHandlerFactory{
		L:           prod.L,
		Srv:         prod.Srv,
		Transaction: ucTransaction.Create(),
	}
	srvTransaction.Create()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		prod.L.Error(fmt.Errorf("failed to listen: %v", err))
	}
	go func() {
		prod.L.Info(fmt.Sprintf("Serving gRPC on %s", lis.Addr()))
		prod.L.Fatal(prod.Srv.Serve(lis))
	}()

	err = transactionPb.RegisterTransactionServiceHandlerFromEndpoint(ctx, muxHttp, lis.Addr().String(), opts)
	if err != nil {
		prod.L.Error(fmt.Errorf("failed to register endpoint transaction: %v", err))
	}
}
