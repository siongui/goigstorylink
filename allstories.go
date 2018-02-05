package igstory

// Get all stories of users that a user follows
// The response will return all users with unexpired stories,
// but only stories of users of unread stories will be returned.

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const UrlAllStories = `https://i.instagram.com/api/v1/feed/reels_tray/`

type RawReelsTray struct {
	Trays []Tray `json:"tray"`
}

type Tray struct {
	Id              int        `json:"id"`
	LatestReelMedia int        `json:"latest_reel_media"`
	User            TrayUser   `json:"user"`
	Items           []TrayItem `json:"items"`
}

type TrayUser struct {
	Username string `json:"username"`
}

type TrayItem struct {
	TakenAt         int64                  `json:"taken_at"`
	DeviceTimestamp int64                  `json:"device_timestamp"` // not reliable value
	ImageVersions2  TrayItemImageVersion2  `json:"image_versions2"`
	VideoVersions   []TrayItemVideoVersion `json:"video_versions"`
	//HasAudio        bool                    `json:"has_audio"`
}

type TrayItemVideoVersion struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

type TrayItemImageVersion2 struct {
	Candidates []struct {
		Url string `json:"url"`
	} `json:"candidates"`
}

func GetAllStories(cfg map[string]string) (trays []Tray, err error) {
	req, err := http.NewRequest("GET", UrlAllStories, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: cfg["ds_user_id"]})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: cfg["sessionid"]})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: cfg["csrftoken"]})

	req.Header.Set("User-Agent", cfg["User-Agent"])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	dec := json.NewDecoder(resp.Body)
	t := RawReelsTray{}
	if err = dec.Decode(&t); err != nil {
		return
	}
	trays = t.Trays

	return
}

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
