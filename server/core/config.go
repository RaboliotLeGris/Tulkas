package core

import (
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type internalConfig struct {
	Dev_mode         bool          `env:"DEV_MODE" envDefault:"false"`
	Log_level        string        `env:"LOG_LEVEL" envDefault:"info"`
	Port             uint32        `env:"PORT" envDefault:"7777"`
	Port_CMD         uint32        `env:"PORT_CMD" envDefault:"8888"`
	Retention_time   time.Duration `env:"DURATION" envDefault:"72h"`
	Storage_folder   string        `env:"STORAGE_FOLDER" envDefault:"/tmp/tulkas/data"`
	UserName         string        `env:"USER_NAME" envDefault:"root"`
	UserHashPassword string        `env:"USER_HASH_PASSWORD,notEmpty"`
}

type Config struct {
	Dev_mode         bool
	Log_level        log.Level
	Port_HTTP        uint32
	Port_CMD         uint32
	Retention_time   time.Duration
	Storage_folder   string
	UserName         string
	UserHashPassword string
}

func NewConfig() (*Config, error) {
	var err error

	internalCfg := internalConfig{}
	if err = env.Parse(&internalCfg); err != nil {
		return nil, err
	}
	log.Debug(internalCfg)

	cfg := Config{}
	cfg.Dev_mode = internalCfg.Dev_mode
	if cfg.Log_level, err = log.ParseLevel(internalCfg.Log_level); err != nil {
		return nil, err
	}
	cfg.Port_HTTP = internalCfg.Port
	cfg.Port_CMD = internalCfg.Port_CMD
	cfg.Retention_time = internalCfg.Retention_time
	cfg.Storage_folder = internalCfg.Storage_folder
	cfg.UserName = internalCfg.UserName
	cfg.UserHashPassword = internalCfg.UserHashPassword

	return &cfg, nil
}
