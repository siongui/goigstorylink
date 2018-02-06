package igstory

// Get all stories of users that a user follows
// The response will return all users with unexpired stories,
// but only stories of users of unread stories will be returned.

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const UrlAllStories = `https://i.instagram.com/api/v1/feed/reels_tray/`

// Used to decode JSON returned by Instagram story API.
type RawReelsTray struct {
	Trays []Tray `json:"tray"`
}

// Used to decode JSON returned by Instagram story API.
type Tray struct {
	Id              int64      `json:"id"`
	LatestReelMedia int64      `json:"latest_reel_media"`
	User            TrayUser   `json:"user"`
	Items           []TrayItem `json:"items"`
}

// Used to decode JSON returned by Instagram story API.
type TrayUser struct {
	Username string `json:"username"`
}

// Used to decode JSON returned by Instagram story API.
type TrayItem struct {
	TakenAt         int64                  `json:"taken_at"`
	DeviceTimestamp int64                  `json:"device_timestamp"` // not reliable value
	ImageVersions2  TrayItemImageVersion2  `json:"image_versions2"`
	VideoVersions   []TrayItemVideoVersion `json:"video_versions"`
	//HasAudio        bool                    `json:"has_audio"`
}

// Used to decode JSON returned by Instagram story API.
type TrayItemVideoVersion struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

// Used to decode JSON returned by Instagram story API.
type TrayItemImageVersion2 struct {
	Candidates []struct {
		Url string `json:"url"`
	} `json:"candidates"`
}

func GetReelsTray() (trays []Tray, err error) {
	req, err := http.NewRequest("GET", UrlAllStories, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: config["ds_user_id"]})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: config["sessionid"]})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: config["csrftoken"]})

	req.Header.Set("User-Agent", config["User-Agent"])

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
