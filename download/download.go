package igstorydl

import (
	"fmt"
	"github.com/siongui/goigstorylink"
	"os"
)

func Download() {
	users, err := igstory.GetUnreadStories()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		for _, story := range user.Stories {
			p := BuildOutputPath(user.Username, story.Url, story.Timestamp)
			// check if file exist
			if _, err = os.Stat(p); os.IsNotExist(err) {
				// file not exists
				fmt.Println(p)
				// run shell wget URL -o path
			}
		}
	}
}
