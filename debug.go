package igstory

import (
	"fmt"
	"time"
)

func PrintItems(items []TrayItem) {
	// One item represents one story
	for _, item := range items {
		// print timestamp of story
		// DO NOT use DeviceTimestamp. It's not reliable
		t := time.Unix(item.TakenAt, 0)
		fmt.Println(t.Format(time.RFC3339))

		// check if the story is video or image
		if len(item.VideoVersions) > 0 {
			// the story is video
			fmt.Println(item.VideoVersions[0].Url)
		} else {
			// the story is image
			fmt.Println(item.ImageVersions2.Candidates[0].Url)
		}
	}
}

func PrintTray(tray Tray) {
	fmt.Print(tray.Id)
	fmt.Print(" : ")
	fmt.Println(tray.User.Username)
	PrintItems(tray.Items)
}

func PrintHighlightTray(tray HighlightTray) {
	fmt.Print(tray.Id)
	fmt.Print(" : ")
	fmt.Println(tray.User.Username)
	fmt.Println(tray.Title)
	PrintItems(tray.Items)
}

func PrintIGUsers(users []IGUser) {
	for _, user := range users {
		fmt.Print(user.Id)
		fmt.Print(" : ")
		fmt.Println(user.Username)
		for _, story := range user.Stories {
			tt := time.Unix(story.Timestamp, 0)
			fmt.Print(tt.Format(time.RFC3339))
			fmt.Print("\t")
			fmt.Println(story.Url)
		}
	}

}
