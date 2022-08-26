package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sanzharanarbay/repository-service-pattern/routes"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")

	prefix := os.Getenv("ROUTE_PREFIX")
	fmt.Println("Server started at " + port + "...")

	router := gin.New()
	routes.ApiRoutes(prefix, router)
	router.Run(host + ":" + port)
}
