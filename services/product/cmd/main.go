package cmd

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/internal/delivery"
	"github.com/YogiTan00/Reseller/services/product/internal/repository"
	"github.com/YogiTan00/Reseller/services/product/internal/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
)

type ProductInit struct {
	Db  *gorm.DB
	Srv *grpc.Server
	L   logger.Logger
}

func (prod *ProductInit) Create(ctx context.Context, muxHttp *runtime.ServeMux, port string, opts []grpc.DialOption) {
	prod.L.EndPoint = "Product"
	repoProduct := &repository.ProductRepositoryFactory{
		Db: *prod.Db,
	}
	ucProduct := &usecase.ProductUsecaseFactory{
		L:           prod.L,
		RepoProduct: repoProduct.Create(),
	}
	srvProduct := &delivery.ProductHandlerFactory{
		L:       prod.L,
		Srv:     prod.Srv,
		Product: ucProduct.Create(),
	}
	srvProduct.Create()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		prod.L.Error(fmt.Errorf("failed to listen: %v", err))
	}
	go func() {
		prod.L.Info(fmt.Sprintf("Serving gRPC on %s", lis.Addr()))
		prod.L.Fatal(prod.Srv.Serve(lis))
	}()

	err = productPb.RegisterProductServiceHandlerFromEndpoint(ctx, muxHttp, lis.Addr().String(), opts)
	if err != nil {
		prod.L.Error(fmt.Errorf("failed to register endpoint product: %v", err))
	}
}
