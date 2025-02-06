package main

import(
	"fmt"
	"path/filepath"
	"hls-server/config"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(filepath.Join("static", "html", "*"))
	r.Static("/static", filepath.Join("static"))

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("main(): cannot load env file")
	}

	// configure routes
	config.SetRoutes(r)
	
	r.Run(":3000")
}