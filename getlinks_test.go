package igstory

import (
	"os"
	"strconv"
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

func ExampleGetUserHighlightStories(t *testing.T) {
	SetUserId(os.Getenv("IG_DS_USER_ID"))
	SetSessionId(os.Getenv("IG_SESSIONID"))
	SetCsrfToken(os.Getenv("IG_CSRFTOKEN"))
	// id of user *instagram* is 25025320
	testid, err := strconv.ParseInt(os.Getenv("IG_TEST_USER_ID"), 10, 64)
	if err != nil {
		t.Error(err)
		return
	}
	users, err := GetUserHighlightStories(testid)
	if err != nil {
		t.Error(err)
		return
	}
	PrintIGUsers(users)
}
