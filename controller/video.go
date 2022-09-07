package controller

import (
	"video-commentary/model"
	"video-commentary/repo"
)

type VideoController struct {
	VideoRepo *repo.VideoRepo
}

func NewVideoController(vr *repo.VideoRepo) *VideoController {
	return &VideoController{
		VideoRepo: vr,
	}
}

func (vc *VideoController) AddVideo(v *model.Video) error {
	vc.VideoRepo.AddVideo(v)

	return nil
}

func (vc *VideoController) GetVideos() ([]model.Video, error) {
	return vc.VideoRepo.GetVideos(), nil
}

func (vc *VideoController) GetVideosByIndex(index int) (*model.Video, error) {
	return vc.VideoRepo.GetVideoByIndex(index), nil
}
