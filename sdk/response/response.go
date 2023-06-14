package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}, msg string) {
	result(c, http.StatusOK, http.StatusOK, data, msg)
}
func Fail(c *gin.Context, data interface{}, msg string) {
	result(c, http.StatusOK, http.StatusBadRequest, data, msg)
}

func result(c *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg})
}
