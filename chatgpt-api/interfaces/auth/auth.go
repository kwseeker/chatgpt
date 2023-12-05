package auth

import (
	"github.com/gin-gonic/gin"
	jwt2 "kwseeker.top/chatgpt/chatgpt-api/domain/auth/jwt"
	user2 "kwseeker.top/chatgpt/chatgpt-api/domain/auth/user"
	"kwseeker.top/chatgpt/chatgpt-api/domain/http/response"
	"kwseeker.top/chatgpt/chatgpt-api/interfaces/router"
	"net/http"
)

// Auth 认证授权相关HTTP接口

func Initialize() {
	//登录，用户名密码换token
	router.Register("/login", http.MethodPost, func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		if username != "kwseeker" {
			response.FailureUserNotExist.Resp(context)
			return
		}
		if password != "123456" {
			response.FailurePasswordUnmatched.Resp(context)
		}
		//实际应该是搜索数据库，获取用户信息
		user := &user2.User{
			Name: "kwseeker", Role: "Admin", Age: 18, Desc: "A man",
		}
		token, err := jwt2.Eng.GenerateWithUser(user)
		if err != nil {
			response.Failure.SetMessage("生成JWT令牌失败").Resp(context)
		}
		response.Success.SetMessage("登录成功").SetData(jwt2.Token{Content: token}).Resp(context)
	})

	//router.Register("/authenticate", http.MethodPost, func(context *gin.Context) {
	//	//认证, 验证token
	//
	//})
}
