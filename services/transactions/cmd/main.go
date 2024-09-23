package cmd

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/config"
	"github.com/YogiTan00/Reseller/pkg/logger"
	transactionPb "github.com/YogiTan00/Reseller/proto/_generated/transaction"
	"github.com/YogiTan00/Reseller/services/transactions/internal/delivery"
	"github.com/YogiTan00/Reseller/services/transactions/internal/repository"
	"github.com/YogiTan00/Reseller/services/transactions/internal/services/product"
	"github.com/YogiTan00/Reseller/services/transactions/internal/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

type TransactionInit struct {
	L   logger.Logger
	Cfg *config.Config
	Db  *gorm.DB
	Srv *grpc.Server
}

func (trs *TransactionInit) Create(ctx context.Context, muxHttp *runtime.ServeMux, port string, opts []grpc.DialOption) {
	trs.L.EndPoint = "Transaction"
	//Grpc Product
	connProd, err := grpc.NewClient(trs.Cfg.PortProduct, opts[0])
	prod := product.ProductServiceFactory{
		L:    trs.L,
		Conn: connProd,
	}
	//Transaction
	repoTransaction := &repository.TransactionRepositoryFactory{
		Db: *trs.Db,
	}
	ucTransaction := &usecase.TransactionUsecaseFactory{
		L:               trs.L,
		RepoTransaction: repoTransaction.Create(),
		ServiceProduct:  prod.Create(),
	}
	srvTransaction := &delivery.TransactionHandlerFactory{
		L:           trs.L,
		Srv:         trs.Srv,
		Transaction: ucTransaction.Create(),
	}
	srvTransaction.Create()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		trs.L.Error(fmt.Errorf("failed to listen: %v", err))
	}
	go func() {
		trs.L.Info(fmt.Sprintf("Serving gRPC on %s", lis.Addr()))
		trs.L.Fatal(trs.Srv.Serve(lis))
	}()

	err = transactionPb.RegisterTransactionServiceHandlerFromEndpoint(ctx, muxHttp, lis.Addr().String(), opts)
	if err != nil {
		trs.L.Error(fmt.Errorf("failed to register endpoint transaction: %v", err))
	}
}
