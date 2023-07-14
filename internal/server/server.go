package server

import (
	"github.com/gin-blog/faye-wong/configs"
	"github.com/gin-blog/faye-wong/internal/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	api *api.Apis
}

func NewServer(apis *api.Apis) *Server {
	s := &Server{
		api: apis,
	}
	return s
}

func (s *Server) StartServe() {
	ginNew := gin.New()
	s.api.SetupRoute(ginNew)
	if err := ginNew.Run(configs.ListenAddr); err != nil {
		return
	}
}
