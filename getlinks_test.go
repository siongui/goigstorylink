package igstory

import (
	"testing"
)

func ExampleGetAllStories(t *testing.T) {
	users, err := GetUnreadStories()
	if err != nil {
		t.Error(err)
		return
	}
	printIGUsers(users)

	if len(users) > 1 {
		tray, err := GetUserStories(users[1].Id, config)
		if err != nil {
			t.Error(err)
			return
		}
		printTray(tray)
	}
}
