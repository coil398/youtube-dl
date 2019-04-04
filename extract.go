package main

import (
	"regexp"
)

// VideoID extract id of video from video url.
func VideoID(url string) string {
	r, _ := regexp.Compile(`(?:v=|\/)([0-9A-Za-z_-]{11}.*)`)
	res := r.FindAllStringSubmatch(url, -1)[0]
	id := res[1]
	return id
}

// WatchURL create a sanitized url.
func WatchURL(id string) string {
	baseURL := "https://www.youtube.com/watch?v="
	watchURL := baseURL + id
	return watchURL
}

// EmbedURL create a embed url.
func EmbedURL(id string) string {
	baseURL := "https://www.youtube.com/embed/"
	embedURL := baseURL + id
	return embedURL
}
