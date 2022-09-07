package main

import (
	"video-commentary/controller"
	"video-commentary/handler"
	"video-commentary/repo"
	"video-commentary/router"
)

func main() {
	vr := repo.NewVideoRepo()
	vc := controller.NewVideoController(vr)
	vh := handler.NewVideoHandler(vc)

	r := router.NewRouter(vh)

	r.Run()
}
