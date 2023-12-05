package router

import (
	"github.com/gin-gonic/gin"
)

// Map 路由路径 -> 请求类型 -> Router
type Map map[string]map[string]*Router

type Router struct {
	Path   string
	Method string //取值 http.method.go
	Func   gin.HandlerFunc
}

// Routers 路由列表
var Routers = make(Map)

func Register(path string, method string, handlerFunc gin.HandlerFunc) {
	router := &Router{
		Path:   path,
		Method: method,
		Func:   handlerFunc,
	}
	if Routers[path] == nil {
		Routers[path] = make(map[string]*Router)
	}
	Routers[path][method] = router
}
