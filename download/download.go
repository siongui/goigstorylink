package igstorydl

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/siongui/goigstorylink"
	"os"
	"os/exec"
)

func Download() {
	users, err := igstory.GetUnreadStories()
	if err != nil {
		panic(err)
	}

	cc := color.New(color.FgCyan)
	rc := color.New(color.FgRed)
	for _, user := range users {
		for _, story := range user.Stories {
			p := BuildOutputPath(user.Username, story.Url, story.Timestamp)
			// check if file exist
			if _, err = os.Stat(p); os.IsNotExist(err) {
				// file not exists
				fmt.Print("Username: ")
				rc.Println(user.Username)
				fmt.Print("Timestamp: ")
				rc.Println(FormatTimestamp(story.Timestamp))
				fmt.Print("Download ")
				cc.Print(story.Url)
				fmt.Print(" to ")
				rc.Print(p)
				fmt.Println(" ...")
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
