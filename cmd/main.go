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

func main() {
	engine := gin.Default()

	restServer := server.NewRestServer(engine)
	restServer.Start()
	

	err := engine.Run()
	if err != nil {
		fmt.Println("Error starting server: " + err.Error())
	} 
}
