package handler

import (
	"net/http"

	"github.com/1024casts/snake/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Response api的返回结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse 返回json
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// GetUserID 返回用户id
func GetUserID(c *gin.Context) uint64 {
	if c == nil {
		return 0
	}

	// uid 必须和 middleware/auth 中的 uid 命名一致
	if v, exists := c.Get("uid"); exists {
		uid, ok := v.(uint64)
		if !ok {
			return 0
		}

		return uid
	}
	return 0
}
