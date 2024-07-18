package main

import (
	"github.com/YogiTan00/Reseller/config"
	"github.com/YogiTan00/Reseller/config/database/mysql"
)

func main() {
	cfg := config.NewConfig()
	_ = mysql.InitMysqlDB(cfg)
	mysql.NewMigration(cfg)

}
