package common

import (
	"fmt"

	"github.com/kalamet/jdoc/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// mysql数据库
func InitDB() *gorm.DB {
	//dbUser :=
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.database"),
		viper.GetString("database.charset"))
	//初始化数据库
	db, ok := gorm.Open(mysql.Open(dsn))
	if ok != nil {
		panic("数据库连接失败, error:" + ok.Error())
	}
	//创建数据库
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
