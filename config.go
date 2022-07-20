package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GormEngine      string `mapstructure:"GORM_ENGINE"`
	GormConnection  string `mapstructure:"GORM_CONNECTION"`
	AppVersion      string `mapstructure:"APP_VERSION"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	KuCoinApiKey    string `mapstructure:"KUCOIN_KEY"`
	KuCoinApiSecret string `mapstructure:"KUCOIN_SECRET"`
	KuCoinApiPass   string `mapstructure:"KUCOIN_PASSPHRASE"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	// Read file path
	viper.AddConfigPath(".")
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
