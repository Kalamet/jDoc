package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kalamet/jdoc/common"
	"github.com/kalamet/jdoc/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		token := cxt.GetHeader("Authorization")
		log.Printf("The token value is: %v", token)
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			cxt.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			cxt.Abort()
			return
		}

		//解析token
		token = token[7:]
		log.Printf("The token2 value is: %v", token)
		t, claims, err := common.ParseToken(token)
		if err != nil || !t.Valid {
			cxt.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			cxt.Abort()
			return
		}

		//判断用户是否存在
		userId := claims.UserId
		log.Printf("The userId is: %v", userId)
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			cxt.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足",
			})
			cxt.Abort()
			return
		}
		//将用户信息传入上下文
		cxt.Set("user", user)
		cxt.Next()
	}

}
