package igstory

import (
	"testing"
)

func ExampleGetUserHighlightStoriesTray(t *testing.T) {
	SetUserId("your user id")
	SetSessionId("your session id")
	SetCsrfToken("your csrf token")
	trays, err := GetUserHighlightStoriesTray(25025320)
	if err != nil {
		t.Error(err)
		return
	}
	for _, tray := range trays {
		PrintHighlightTray(tray)
	}
}
