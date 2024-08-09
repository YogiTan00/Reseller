package main

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/config"
	initMysql "github.com/YogiTan00/Reseller/config/database/mysql"
	"github.com/YogiTan00/Reseller/config/html"
	productPb "github.com/YogiTan00/Reseller/proto/_generated/product"
	"github.com/YogiTan00/Reseller/services/product/cmd"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"net/http"
)

func main() {
	var (
		l       = logrus.New()
		cfg     = config.NewConfig()
		db      = initMysql.InitMysqlDB(cfg)
		ctx     = context.Background()
		muxHttp = runtime.NewServeMux()
		opts    = []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
	)
	l.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	initMysql.NewMigration(cfg)
	connDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		l.Error(err)
	}
	srv := grpc.NewServer()
	product := cmd.ProductHandlerFactory{
		Db:  connDb,
		Srv: srv,
	}
	product.Create()
	lis, err := net.Listen("tcp", cfg.PortProduct)
	if err != nil {
		l.Error(fmt.Errorf("failed to listen: %v", err))
	}
	go func() {
		l.Info(fmt.Sprintf("Serving gRPC on %s", lis.Addr()))
		l.Fatal(srv.Serve(lis))
	}()
	err = productPb.RegisterProductServiceHandlerFromEndpoint(ctx, muxHttp, lis.Addr().String(), opts)
	if err != nil {
		l.Error(fmt.Errorf("failed to register endpoint product: %v", err))
	}

	r := mux.NewRouter()
	r.HandleFunc("/", html.HomeHandler())
	r.HandleFunc("/product", html.ProductHandler())
	r.PathPrefix("/").Handler(muxHttp)

	l.Info(fmt.Sprintf("Serving HTTP and gRPC-Gateway on %s", cfg.Port))
	l.Fatal(http.ListenAndServe(cfg.Port, r))
}
