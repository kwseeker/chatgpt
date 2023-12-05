package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"kwseeker.top/chatgpt/chatgpt-api/interfaces/router"
	"net/http"
)

const (
	Port = ":8080"
)

type Server struct {
	Engine  *gin.Engine
	Addr    string
	routers router.Map
}

func New() *Server {
	s := &Server{
		Engine: gin.Default(),
		Addr:   Port,
	}
	s.routers = router.Routers
	return s
}

func (s *Server) Start() {
	if s.Engine == nil {
		s.Engine = gin.Default()
	}

	//遍历路由列表注册到Gin引擎
	s.setupRouters()

	//启动Gin引擎
	err := s.Engine.Run(s.Addr)
	if err != nil {
		logrus.Fatal("server start failed, err=", err)
	}
}

func (s *Server) setupRouters() {
	for _, method2Routers := range s.routers {
		for _, r := range method2Routers {
			switch r.Method {
			case http.MethodGet:
				s.Engine.GET(r.Path, r.Func)
			case http.MethodHead:
				s.Engine.HEAD(r.Path, r.Func)
			case http.MethodPost:
				s.Engine.POST(r.Path, r.Func)
			case http.MethodPut:
				s.Engine.PUT(r.Path, r.Func)
			case http.MethodPatch:
				s.Engine.PATCH(r.Path, r.Func)
			case http.MethodDelete:
				s.Engine.DELETE(r.Path, r.Func)
			case http.MethodOptions:
				s.Engine.OPTIONS(r.Path, r.Func)
			default:
				s.Engine.NoMethod(r.Func)
			}
		}
	}
}
