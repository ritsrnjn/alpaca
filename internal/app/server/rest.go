package server

import (
	"github.com/gin-gonic/gin"
)

type restServer struct {
	engine *gin.Engine
}

func NewRestServer(engine *gin.Engine) *restServer {
	return &restServer{
		engine: engine,
	}
}

func (s *restServer) Start() {
	s.engine.GET("/random", func(c *gin.Context) {
		c.String(200, "You have reached the random endpoint running on alpaca server")
	},
	)
}