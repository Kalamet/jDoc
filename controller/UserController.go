package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kalamet/jdoc/common"
	"github.com/kalamet/jdoc/dto"
	"github.com/kalamet/jdoc/model"
	"github.com/kalamet/jdoc/response"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var db = common.GetDB()
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	//检查参数
	//检查phone是否注册
	if model.IsRegister(db, phone) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "用户已存在",
		})
		return
	}
	//创建用户
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := model.User{
		Phone:         phone,
		Password:      string(hashPassword),
		LastLoginTime: time.Now().Unix(),
	}

	db.Create(&newUser)
	token, _ := common.ReleaseToken(newUser)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
		"data": gin.H{
			"userId": newUser.ID,
			"token":  token,
		},
	})
}

func Login(cxt *gin.Context) {
	phone := cxt.PostForm("phone")
	password := cxt.PostForm("password")

	//判断用户是否存在
	var db = common.GetDB()
	var user model.User
	if !model.IsRegister(db, phone) {
		cxt.JSON(200, gin.H{
			"message": "用户不存在",
		})
		return
	}
	//判断密码是否正确
	db.Where("phone = ?", phone).First(&user)
	fmt.Println(user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		cxt.JSON(200, gin.H{
			"message": "密码错误，请重新输入",
		})
		return
	}
	//返回用户token
	token, err := common.ReleaseToken(user)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{
			"message": "系统错误",
		})
		log.Printf("Release token err: %v", err)
		return
	}
	cxt.JSON(200, gin.H{
		"message": "ok",
		"userId":  user.ID,
		"token":   token,
	})
}

func UserInfo(cxt *gin.Context) {
	user, _ := cxt.Get("user")
	response.Json(cxt, http.StatusOK, 0, "ok", gin.H{"user": dto.ToUserDto(user.(model.User))})
}
