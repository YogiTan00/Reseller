package config

import (
	"github.com/YogiTan00/Reseller/pkg/utils"
	"os"
)

type Config struct {
	Port        string
	PortProduct string
	Domain      string
	Debug       string
	Address     string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPass      string
	DbName      string
	Migration   string
	PathLogs    string
}

func NewConfig() *Config {
	return &Config{
		Port:        utils.NewEnv("port"),
		PortProduct: utils.NewEnv("port_product"),
		Domain:      utils.NewEnv("domain"),
		Debug:       utils.NewEnv("debug"),
		Address:     utils.NewEnv("address"),
		DbHost:      utils.NewEnv("db_host"),
		DbPort:      utils.NewEnv("db_port"),
		DbUser:      utils.NewEnv("db_user"),
		DbPass:      utils.NewEnv("db_pass"),
		DbName:      utils.NewEnv("db_name"),
		Migration:   utils.NewEnv("migration"),
		PathLogs:    utils.NewEnvDefault("path_logs", os.TempDir()),
	}
}
