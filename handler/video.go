package handler

import (
	"net/http"
	"video-commentary/controller"
	"video-commentary/model"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	Controller *controller.VideoController
}

func NewVideoHandler(c *controller.VideoController) *VideoHandler {
	return &VideoHandler{
		Controller: c,
	}
}

func (h *VideoHandler) PostVideo(ctx *gin.Context) {
	var input model.Video

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding JSON",
		})

		return
	}

	if err := h.Controller.AddVideo(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "succeed post video",
	})
}

func (h *VideoHandler) GetVideos(ctx *gin.Context) {
	if videos, err := h.Controller.GetVideos(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   videos,
		})
	}
}
