package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"kwseeker.top/chatgpt/chatgpt-api/domain/http/response"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/success", func(c *gin.Context) {
		c.String(http.StatusOK, "test success by kwseeker")
	})

	r.GET("/verify", func(c *gin.Context) {
		token := c.Query("token")
		logrus.Infof("验证 token: %s", token)
		if token == "success" {
			c.JSON(http.StatusOK, &response.Entity{
				Code:    http.StatusOK,
				Status:  "success",
				Message: "请求成功",
			})
			return
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		logrus.Errorln("start server failed on port 8080", err)
	}
}
