package nets

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义统一响应结构体
type Response struct {
	Code int32       `json:"code"` // 业务状态码，非 HTTP 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据（可以是数组、对象、字符串等）
	Time int64       `json:"time"` // 响应时间戳
}

// 封装成功响应函数
func Success(c *gin.Context, data interface{}) {
	resp := Response{
		Code: 200,
		Msg:  "success",
		Data: data,
		Time: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, resp)
}

// 封装失败响应函数
func Fail(c *gin.Context, code int32, msg string) {
	resp := Response{
		Code: code,
		Msg:  msg,
		Data: nil, // 失败时数据为空
		Time: time.Now().Unix(),
	}
	c.JSON(http.StatusOK, resp)
}
