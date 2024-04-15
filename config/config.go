package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

var (
	cfg  Configure
	once sync.Once
)

func GetConfigure() *Configure {
	once.Do(func() {
		if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
			panic(err)
		}
	})
	return &cfg
}

type Configure struct {
	Title string `toml:"title"`
	Proto struct {
		TCP  string `toml:"tcp"`
		Ws   string `toml:"ws"`
		Mqtt string `toml:"mqtt"`
	}
	Mqtt struct {
		Addr     string `toml:"addr"`
		UserName string `toml:"username"`
		Password string `toml:"password"`
		Expire   int    `toml:"expire"`
	}
	Database struct {
		Type string `toml:"type"`
		// DSN  struct {
		Host         string `toml:"host"`
		DB           string `toml:"db" `
		Username     string `toml:"username" `
		Password     string `toml:"password"`
		Port         int    `toml:"port"`
		MaxIdleConns int    `toml:"max_idle_conns"`
		MaxOpenConns int    `toml:"max_open_conns"`
		// } `toml:"dsn"`
	} `toml:"database"`
	Kafka struct {
		// Offset        string `toml:"offset" `
		// SaslMechanism string `toml:"sasl_mechanism" envconfig:"SMCB_APP_BACKEND_KAFKA_SASL_MECHANISM"`
		// SaslUser      string `toml:"sasl_user" envconfig:"SMCB_APP_BACKEND_KAFKA_SASL_USER"`
		// SaslPassword  string `toml:"sasl_password" envconfig:"SMCB_APP_BACKEND_KAFKA_SASL_PASSWORD"`
		Addr string `toml:"addr"`
		Port string `toml:"port"`
		// PoolSize      int    `toml:"pool_size" envconfig:"SMCB_APP_BACKEND_KAFKA_POOL_SIZE"`
		// SaslEnable    bool   `toml:"sasl_enable" envconfig:"SMCB_APP_BACKEND_KAFKA_SASL_ENABLE"`
	} `toml:"kafka"`
}
