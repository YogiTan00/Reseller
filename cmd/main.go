package main

import (
	"context"
	"fmt"
	"github.com/YogiTan00/Reseller/config"
	initMysql "github.com/YogiTan00/Reseller/config/database/mysql"
	"github.com/YogiTan00/Reseller/config/html"
	"github.com/YogiTan00/Reseller/pkg/logger"
	cmdProd "github.com/YogiTan00/Reseller/services/product/cmd"
	cmdTrans "github.com/YogiTan00/Reseller/services/transactions/cmd"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	var (
		l = logger.Logger{
			EndPoint: "Main",
		}
		cfg     = config.NewConfig()
		db      = initMysql.InitMysqlDB(cfg)
		ctx     = context.Background()
		muxHttp = runtime.NewServeMux()
		opts    = []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}
	)
	l.CreateNewLog()
	initMysql.NewMigration(cfg)
	connDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		l.Error(err)
	}
	srv := grpc.NewServer()

	product := cmdProd.ProductInit{
		Db:  connDb,
		Srv: srv,
		L:   l,
	}
	product.Create(ctx, muxHttp, cfg.PortProduct, opts)
	transaction := cmdTrans.TransactionInit{
		L:   l,
		Cfg: cfg,
		Db:  connDb,
		Srv: srv,
	}
	transaction.Create(ctx, muxHttp, cfg.PortTransaction, opts)

	r := mux.NewRouter()
	r.HandleFunc("/", html.HomeHandler())
	r.HandleFunc("/product", html.ProductHandler())
	r.PathPrefix("/").Handler(muxHttp)

	l.Info(fmt.Sprintf("Serving HTTP and gRPC-Gateway on %s", cfg.Port))
	l.Fatal(http.ListenAndServe(cfg.Port, r))
}
