package main

import (
	"flag"
	"fmt"
)

type StreamMonoState struct {
}

type Video struct {
	url             string
	id              string
	watchURL        string
	embedURL        string
	streamMonoState StreamMonoState
}

// prefetch gets all necessary data.
func prefetch() {

}

func initializer() {

}

func main() {
	flag.Parse()
	url := flag.Arg(0)
	video := new(Video)
	video.url = url
	video.id = VideoID(url)
	video.watchURL = WatchURL(video.id)
	video.embedURL = EmbedURL(video.id)
	fmt.Println(video)
	prefetch()
	initializer()
}
