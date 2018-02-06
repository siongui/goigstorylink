package igstorydl

import (
	"fmt"
	"github.com/siongui/goigstorylink"
	"os"
	"os/exec"
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
				fmt.Println("Download " + story.Url + " to " + p + " ...")
				// run shell wget URL -o path
				cmd := exec.Command("wget", story.Url, "-O", p)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
