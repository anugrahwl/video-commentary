package router

import (
	"video-commentary/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(vh *handler.VideoHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/videos", vh.PostVideo)
	r.GET("/videos", vh.GetVideos)

	return r
}
