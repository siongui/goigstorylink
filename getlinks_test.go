package igstory

import (
	"testing"
)

func ExampleGetAllStories(t *testing.T) {
	SetUserId("your user id")
	SetSessionId("your session id")
	SetCsrfToken("your csrf token")
	users, err := GetAllStories()
	if err != nil {
		t.Error(err)
		return
	}
	printIGUsers(users)
}
