package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var c *Config

type DatabaseHost struct {
	Hostname string
	Username string
	Password string
	DBName   string
}

type Config struct {
	Database struct {
		Reader DatabaseHost
	}
}

func init() {
	viper.SetConfigName("default")
	currentPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/%s", currentPath, "resources")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	c = new(Config)
	viper.Unmarshal(c)
}

func GetConfig() *Config {
	return c
}
