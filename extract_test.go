package main

import (
	"testing"
)

func TestVideoID(t *testing.T) {
	url := "https://www.youtube.com/watch?v=qcSVPCWCPbk"
	result := VideoID(url)
	if result != "qcSVPCWCPbk" {
		t.Errorf("expecting qcSVPCWCPbk, got %s", result)
	}
}

func TestWatchURL(t *testing.T) {
	id := "qcSVPCWCPbk"
	result := WatchURL(id)
	if result != "https://www.youtube.com/watch?v=qcSVPCWCPbk" {
		t.Errorf("expecting https://www.youtube.com/watch?v=qcSVPCWCPbk, got %s", result)
	}
}

func TestEmbedURL(t *testing.T) {
	id := "qcSVPCWCPbk"
	result := EmbedURL(id)
	if result != "https://www.youtube.com/embed/qcSVPCWCPbk" {
		t.Errorf("expecting https://www.youtube.com/embed/qcSVPCWCPbk, got %s", result)
	}
}
