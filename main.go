package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kalamet/jdoc/common"
	"github.com/spf13/viper"
)

func main() {
	common.InitConfig()
	common.InitDB()
	r := gin.Default()
	r = Routes(r)
	port := viper.GetString("server.port")
	log.Printf("port: %v", port)
	panic(r.Run(":" + port))
}
