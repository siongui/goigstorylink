package igstory

import (
	"testing"
)

func ExampleGetUserHighlightStories(t *testing.T) {
	SetUserId("your user id")
	SetSessionId("your session id")
	SetCsrfToken("your csrf token")
	// Get JSON bytes of highlight stories of user *instagram*
	b, err := GetUserHighlightStories(25025320)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))
}
