package main

import (
	"fmt"
	"github.com/YogiTan00/Reseller/config"
	initMysql "github.com/YogiTan00/Reseller/config/database/mysql"
	"github.com/YogiTan00/Reseller/pkg/logger"
	"github.com/YogiTan00/Reseller/services/product"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
)

func main() {
	l := logger.NewLogger("-=Main=-")
	cfg := config.NewConfig()
	db := initMysql.InitMysqlDB(cfg)
	initMysql.NewMigration(cfg)
	conn, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		l.Error(err)
		return
	}
	srv := grpc.NewServer()
	prod := product.ProductHandlerFactory{
		Db:  *conn,
		Srv: srv,
	}
	prod.Create()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		l.Error(fmt.Errorf("failed to listen: %v", err))
	}
	err = srv.Serve(lis)
	if err != nil {
		l.Error(err)
		return
	}

}
