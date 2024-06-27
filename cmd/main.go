package main

import (
	"Reseller/config"
	"Reseller/config/database/mysql"
)

func main() {
	cfg := config.NewConfig()
	_ = mysql.InitMysqlDB(cfg)
	mysql.NewMigration(cfg)

}
