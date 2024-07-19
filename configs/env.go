package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	RPCUrl string `mapstructure:"RPC_URL"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Could not read config file: %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
