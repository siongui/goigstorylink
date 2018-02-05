package igstory

import (
	"testing"
	"time"
)

func TestGetAllStories(t *testing.T) {
	users, err := GetUnreadStories()
	if err != nil {
		t.Error(err)
		return
	}

	for _, user := range users {
		t.Log(user.Id)
		t.Log(user.Username)
		for _, story := range user.Stories {
			tt := time.Unix(story.Timestamp, 0)
			t.Log(tt.Format(time.RFC3339))
			t.Log(story.Url)
		}
	}

	if len(users) > 1 {
		tray, err := GetUserStories(users[1].Id, config)
		if err != nil {
			t.Error(err)
			return
		}
		printTray(tray)
	}
}
