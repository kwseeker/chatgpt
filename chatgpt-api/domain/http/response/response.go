package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Success                  = New(0, "success", "请求成功", nil)
	Failure                  = New(1, "failure", "请求失败", nil)
	FailureUserNotExist      = New(1001, "fail_user_not_exist", "用户不存在", nil)
	FailurePasswordUnmatched = New(1002, "fail_password_wrong", "密码错误", nil)
)

type Data interface{}

type Entity struct {
	Code    int
	Status  string
	Message string
	Data    Data
}

func New(code int, status string, message string, data Data) *Entity {
	return &Entity{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (e *Entity) Resp(context *gin.Context) {
	context.JSON(http.StatusOK, e.toGinH())
}

// toGinH
func (e *Entity) toGinH() gin.H {
	return gin.H{
		"code":    e.Code,
		"status":  e.Status,
		"message": e.Message,
		"data":    e.Data,
	}
}

func (e *Entity) SetMessage(message string) *Entity {
	e.Message = message
	return e
}

func (e *Entity) SetData(data Data) *Entity {
	e.Data = data
	return e
}
