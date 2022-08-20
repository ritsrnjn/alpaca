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
	
	s.engine.GET("/play", func(c *gin.Context) {
		c.JSON(200, gin.H{
            "message": "Billu-007 You have reached the base endpoint running on alpaca server.",
        })
	},
	)

	s.engine.GET("/game", func(c *gin.Context) {
		c.JSON(200, gin.H{
            "game": "server",
        })
	},
	)

}