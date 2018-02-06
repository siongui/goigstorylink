package igstorydl

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/siongui/goigstorylink"
	"os"
	"os/exec"
	"time"
)

func Wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PrintDownloadInfo(username, url, filepath string, timestamp int64) {
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

func DownloadIGUser(user igstory.IGUser) {
	for _, story := range user.Stories {
		// BuildOutputFilePath also create dir if not exist
		p := BuildOutputFilePath(user.Username, story.Url, story.Timestamp)
		// check if file exist
		if _, err := os.Stat(p); os.IsNotExist(err) {
			// file not exists
			PrintDownloadInfo(user.Username, story.Url, p, story.Timestamp)
			err = Wget(story.Url, p)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func DownloadUnread() {
	users, err := igstory.GetUnreadStories()
	if err != nil {
		// return error? or just print?
		fmt.Println(err)
		return
	}

	for _, user := range users {
		DownloadIGUser(user)
	}
}

func DownloadAll() {
	users, err := igstory.GetAllStories()
	if err != nil {
		// return error? or just print?
		fmt.Println(err)
		return
	}

	for _, user := range users {
		DownloadIGUser(user)
	}
}

func MonitorAndDownload(userid, sessionid, csrftoken string) {
	igstory.SetUserId(userid)
	igstory.SetSessionId(sessionid)
	igstory.SetCsrfToken(csrftoken)

	cc := color.New(color.FgCyan)
	rc := color.New(color.FgRed)
	sleepInterval := 30 // seconds
	count := 0
	for {
		if count == 0 {
			DownloadAll()
			cc.Println("Download all stories finished")
		} else {
			DownloadUnread()
			cc.Println("Download unread stories finished")
		}
		count++
		count %= 5

		rc.Print(time.Now().Format(time.RFC3339))
		fmt.Print(": sleep ")
		cc.Print(sleepInterval)
		fmt.Println(" second")
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
