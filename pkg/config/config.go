package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `json:"app" mapstructure:"app"`
	Database DatabaseConfig `json:"database" mapstructure:"database"`
}

type AppConfig struct {
	Name string `json:"name" mapstructure:"name"`
	Port int32  `json:"port" mapstructure:"port"`
}

type DatabaseConfig struct {
	Host               string `json:"host" mapstructure:"host"`
	Username           string `json:"username" mapstructure:"username"`
	Password           string `json:"password" mapstructure:"password"`
	Port               int32  `json:"port" mapstructure:"port"`
	Name               string `json:"name" mapstructure:"name"`
	Schema             string `json:"schema" mapstructure:"schema"`
	MaxIdleConnections int    `json:"maxIdleConnections" mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections" mapstructure:"maxOpenConnections"`
	DebugMode          bool   `json:"debugMode" mapstructure:"debugMode"`
}

func NewConfig(configPath string) *Config {
	fmt.Println("Load Config.... ")

	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.App.Port)
}
