package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./files/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error read config file: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("error unmarshal config: %v", err)
	}

	// debug only
	tempDebug31, _ := json.Marshal(cfg)
	fmt.Printf("DEBUG :%s \n", string(tempDebug31))

	return cfg
}
