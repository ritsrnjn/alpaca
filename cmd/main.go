package main

//import gin
import (
	"codeingerman/alpaca/internal/app/server"
	
	"fmt"

	"github.com/gin-gonic/gin"
)


const (
	appName = "alpaca"
	appPort = "8080"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	engine := gin.Default()
	engine.Use(CORSMiddleware())

	restServer := server.NewRestServer(engine)
	restServer.Start()
	

	err := engine.Run()
	if err != nil {
		fmt.Println("Error starting server: " + err.Error())
	} 
}
