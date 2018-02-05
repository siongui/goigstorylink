package igstory

import (
	"fmt"
	"time"
)

func printTray(tray Tray) {
	fmt.Print(tray.Id)
	fmt.Print(" : ")
	fmt.Println(tray.User.Username)

	// One item represents one story
	for _, item := range tray.Items {
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
