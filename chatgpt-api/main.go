package main

import (
	"kwseeker.top/chatgpt/chatgpt-api/domain/server"
	"kwseeker.top/chatgpt/chatgpt-api/interfaces/auth"
)

func main() {
	//组件初始化
	auth.Initialize()
	//启动服务器
	s := server.New()
	s.Start()
}
