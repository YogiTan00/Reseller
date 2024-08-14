package config

import (
	"fmt"
	"github.com/YogiTan00/Reseller/pkg/utils"
	"os"
)

const (
	Port            = "port"
	PortProduct     = "port_product"
	PortTransaction = "port_transaction"
	Domain          = "domain"
	Debug           = "debug"
	Address         = "address"
	DbHost          = "db_host"
	DbPort          = "db_port"
	DbUser          = "db_user"
	DbPass          = "db_pass"
	DbName          = "db_name"
	Migration       = "migration"
	PathLogs        = "path_logs"
)

type Config struct {
	Port            string
	PortProduct     string
	PortTransaction string
	Domain          string
	Debug           string
	Address         string
	DbHost          string
	DbPort          string
	DbUser          string
	DbPass          string
	DbName          string
	Migration       string
	PathLogs        string
}

func NewConfig() *Config {
	cfg := &Config{
		Port:            utils.NewEnv(Port),
		PortProduct:     utils.NewEnv(PortProduct),
		PortTransaction: utils.NewEnv(PortTransaction),
		Domain:          utils.NewEnv(Domain),
		Debug:           utils.NewEnv(Debug),
		Address:         utils.NewEnv(Address),
		DbHost:          utils.NewEnv(DbHost),
		DbPort:          utils.NewEnv(DbPort),
		DbUser:          utils.NewEnv(DbUser),
		DbPass:          utils.NewEnv(DbPass),
		DbName:          utils.NewEnv(DbName),
		Migration:       utils.NewEnv(Migration),
		PathLogs:        utils.NewEnvDefault(PathLogs, os.TempDir()),
	}
	err := cfg.Validate()
	if err != nil {
		panic(utils.Color("red", err.Error()))
	}
	return cfg
}

func (cfg *Config) Validate() error {
	required := []string{
		Port,
		PortProduct,
		Domain,
		DbHost,
		DbPort,
		DbUser,
		DbPass,
		DbName,
		Migration,
	}
	for _, v := range required {
		if utils.NewEnv(v) == "" {
			return fmt.Errorf("env %s is required", v)
		}
	}

	return nil
}
