package igstorydl

import (
	"github.com/siongui/goigstorylink"
	"testing"
)

func ExampleDownload(t *testing.T) {
	igstory.SetUserId("")
	igstory.SetSessionId("")
	igstory.SetCsrfToken("")

	Download()
}
