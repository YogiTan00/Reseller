package main

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/config"
	initMysql "github.com/YogiTan00/Reseller/config/database/mysql"
	"github.com/YogiTan00/Reseller/pkg/logger"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/cmd"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"net/http"
)

var (
	l    = logger.NewLogger("-=Main=-")
	cfg  = config.NewConfig()
	db   = initMysql.InitMysqlDB(cfg)
	ctx  = context.Background()
	mux  = runtime.NewServeMux()
	opts = []grpc.DialOption{
		grpc.WithInsecure()}
)

func main() {
	initMysql.NewMigration(cfg)
	connDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		l.Info(err)
	}
	srv := grpc.NewServer()
	product := cmd.ProductHandlerFactory{
		Db:  connDb,
		Srv: srv,
	}
	product.Create()
	lis, err := net.Listen("tcp", cfg.PortProduct)
	if err != nil {
		l.Info(fmt.Errorf("failed to listen: %v", err))
	}
	go func() {
		l.Info(fmt.Sprintf("Serving gRPC on %s", lis.Addr()))
		l.Fatal(srv.Serve(lis))
	}()
	err = productPb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, lis.Addr().String(), opts)
	l.Info(fmt.Sprintf("Serving gRPC-Gateway on %s", cfg.Port))
	l.Fatal(http.ListenAndServe(cfg.Port, mux))
}
