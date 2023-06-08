package common

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	basePath, _ := os.Getwd()
	log.Printf("base path is : %v", basePath)

	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath(basePath + "/config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("viper error: %v", err)
	}
}
