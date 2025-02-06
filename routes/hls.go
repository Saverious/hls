package routes

import (
	"hls-server/handlers"
	"github.com/gin-gonic/gin"
)

func HLS(r *gin.Engine){
	v1 := r.Group("/")
	{
		v1.GET("/", handlers.Home)
	}
}