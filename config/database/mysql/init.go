package mysql

import (
	"Reseller/config"
	"Reseller/pkg/logger"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMysqlDB(cfg config.Config) *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
		log      = logger.NewLogger("Init Mysql")
	)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, errMysql = sql.Open(`mysql`, dsn)

	if errMysql != nil {
		panic(errMysql)
	}

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)
	log.Info("Success Connect to Mysql")
	return dbConn
}

func NewMigration(cfg config.Config) {
	var (
		log = logger.NewLogger("New Migration")
	)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Error(err)
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Error(err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://config/database/mysql/migrations",
		"reseller",
		driver,
	)
	if err != nil {
		log.Error(err)
		return
	}

	mgrt, err := strconv.Atoi(cfg.Migration)
	if err != nil {
		log.Error(err)
		return
	}
	err = m.Steps(mgrt)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return
		}
		log.Error(err)
		return
	}
}
