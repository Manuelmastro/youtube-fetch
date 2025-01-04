package main

import (
	"assignment/config"
	youtubeapi "assignment/yt_api"
)

func main() {
	config.LoadEnv()
	go youtubeapi.BackgroundProcess()
	select {}
}
