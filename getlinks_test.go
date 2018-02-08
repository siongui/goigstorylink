package igstory

import (
	"os"
	"testing"
)

func ExampleGetAllStories(t *testing.T) {
	SetUserId(os.Getenv("IG_DS_USER_ID"))
	SetSessionId(os.Getenv("IG_SESSIONID"))
	SetCsrfToken(os.Getenv("IG_CSRFTOKEN"))
	users, err := GetAllStories()
	if err != nil {
		t.Error(err)
		return
	}
	PrintIGUsers(users)
}

func ExampleGetUnreadStories(t *testing.T) {
	SetUserId(os.Getenv("IG_DS_USER_ID"))
	SetSessionId(os.Getenv("IG_SESSIONID"))
	SetCsrfToken(os.Getenv("IG_CSRFTOKEN"))
	users, err := GetUnreadStories()
	if err != nil {
		t.Error(err)
		return
	}
	PrintIGUsers(users)
}
