package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	AppVersion    string `mapstructure:"APP_VERSION"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func main() {
	// load app.env file data to struct
	config, err := LoadConfig(".")

	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	fmt.Printf(" -----%s----\n", "Reading Environment variables Using Viper package")
	fmt.Printf(" %s = %v \n", "Application_Environment", config.AppEnv)
	// not defined
	fmt.Printf(" %s = %s \n", "DB_DRIVE", config.DBDriver)
	fmt.Printf(" %s = %s \n", "Server_Listening_Address", config.ServerAddress)
	fmt.Printf(" %s = %s \n", "Database_User_Name", config.DBUser)
	fmt.Printf(" %s = %s \n", "Database_User_Password", config.DBPass)

}

func LoadConfig(path string) (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
