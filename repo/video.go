package repo

import "video-commentary/model"

type VideoRepo struct {
	Videos []model.Video
}

func NewVideoRepo() *VideoRepo {
	return &VideoRepo{
		Videos: []model.Video{},
	}
}

func (vr *VideoRepo) AddVideo(v *model.Video) {
	vr.Videos = append(vr.Videos, *v)
}

func (vr *VideoRepo) GetVideoByIndex(index int) *model.Video {
	return &vr.Videos[index]
}

func (vr *VideoRepo) GetVideos() []model.Video {
	return vr.Videos
}
