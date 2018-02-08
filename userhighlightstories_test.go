package igstory

import (
	"testing"
)

func ExampleGetUserHighlightStories(t *testing.T) {
	SetUserId("your user id")
	SetSessionId("your session id")
	SetCsrfToken("your csrf token")
	trays, err := GetUserHighlightStories(25025320)
	if err != nil {
		t.Error(err)
		return
	}
	for _, tray := range trays {
		PrintHighlightTray(tray)
	}
}
