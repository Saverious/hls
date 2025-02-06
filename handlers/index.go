package handlers

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
) 

// load html file
func Home(c *gin.Context){
	endpoint := os.Getenv("SERVER_URL")
	src := fmt.Sprintf("%v%v", endpoint, "/static/media/index.m3")
	fmt.Printf("Home src: %v\n", src)

	c.HTML(200, "index.html", gin.H{
		"title": "Streaming with HLS",
		"heading": "Live streaming a video file",
		"src": src,
	})
}