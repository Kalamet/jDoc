package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Json(cxt *gin.Context, httpStatus int, code int, message string, data gin.H) {
	cxt.JSON(httpStatus, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func Success(cxt *gin.Context, message string, data gin.H) {
	Json(cxt, http.StatusOK, 200, message, data)
}

func Fail(cxt *gin.Context, message string, data gin.H) {
	Json(cxt, http.StatusOK, 400, message, data)
}
