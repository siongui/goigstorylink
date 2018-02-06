package igstorydl

import (
	"testing"
)

func TestStripQueryString(t *testing.T) {
	u := StripQueryString("https://example.com/myvideo.mp3?abc=d")
	if u != "https://example.com/myvideo.mp3" {
		t.Error(u)
	}
}
