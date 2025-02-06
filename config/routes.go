package config

import (
	"hls-server/routes"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine){
	routes.HLS(r)
}