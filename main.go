package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	//延迟关闭数据库
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//注册
	r.POST("register", func(ctx *gin.Context) {
		//获取参数值
		phone := ctx.PostForm("phone")
		password := ctx.PostForm("password")
		//检查参数
		//检查phone是否注册
		if isRegister(db, phone) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    0,
				"message": "用户已存在",
			})
			return
		}
		//创建用户
		newUser := User{
			Phone:         phone,
			Password:      password,
			LastLoginTime: time.Now().Unix(),
		}
		db.Create(&newUser)

		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "注册成功",
			"data": gin.H{
				"phone": phone,
			},
		})
	})
	//登录
	panic(r.Run(":8088"))
}

func isRegister(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	return user.ID != 0
}

func initDB() *gorm.DB {
	//初始化数据库
	db, ok := gorm.Open(mysql.Open(dsn))
	if ok != nil {
		panic("数据库连接失败, error:" + ok.Error())
	}
	//创建数据库
	db.AutoMigrate(&User{})
	return db
}

// mysql数据库
const (
	dbUser     string = "root"
	dbPassword string = "root"
	dbName     string = "jdoc"
	dbPort     string = "3306"
	dbHost     string = "127.0.0.1"
	charset    string = "utf8mb4"
)

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
	dbUser, dbPassword, dbHost, dbPort, dbName, charset)

type Config struct {
	NowFunc func() time.Time
}

// 用户表
type User struct {
	ID            int64  `gorm:"primaryKey"`
	Phone         string `gorm:"type:varchar(20);unique;not null;"`
	Password      string `gorm:"type:varchar(100);not null;default:''"`
	Name          string `gorm:"type:varchar(100)"`
	Avatar        string `gorm:"type:varchar(255)"`
	LastLoginTime int64  `gorm:"autoCreatedTime"`
	CreatedAt     int64  `gorm:"autoCreatedTime"`
	UpdatedAt     int64  `gorm:"autoCreatedTime"`
}
