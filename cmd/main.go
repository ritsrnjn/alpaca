package main

//import gin
import (
	"alpaca/internal/app/server"
	
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)


const (
	appName = "alpaca"
	appPort = "8080"
)

func main() {
	port := os.Getenv("PORT")

	engine := gin.Default()

	restServer := server.NewRestServer(engine)
	restServer.Start()
	

	err := engine.Run(":" + port)
	if err != nil {
		fmt.Println("Error starting server: " + err.Error())
	} 
}
