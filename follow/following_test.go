package igstory

import (
	"os"
	"testing"
)

func ExampleGetFollowing(t *testing.T) {
	rf, err := GetFollowing(
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rf)
}
