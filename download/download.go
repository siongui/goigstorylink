package igstorydl

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/siongui/goigstorylink"
	"os"
	"os/exec"
)

func wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func printDownloadInfo(username, url, filepath string, timestamp int64) {
	cc := color.New(color.FgCyan)
	rc := color.New(color.FgRed)
	fmt.Print("Username: ")
	rc.Println(username)
	fmt.Print("Story timestamp: ")
	rc.Println(FormatTimestamp(timestamp))
	fmt.Print("Download ")
	cc.Print(url)
	fmt.Print(" to ")
	rc.Print(filepath)
	fmt.Println(" ...")
}

func Download() {
	users, err := igstory.GetUnreadStories()
	if err != nil {
		// TODO: return error
		fmt.Println(err)
		return
	}

	for _, user := range users {
		for _, story := range user.Stories {
			p := BuildOutputPath(user.Username, story.Url, story.Timestamp)
			// check if file exist
			if _, err = os.Stat(p); os.IsNotExist(err) {
				// file not exists
				printDownloadInfo(user.Username, story.Url, p, story.Timestamp)
				err = wget(story.Url, p)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
