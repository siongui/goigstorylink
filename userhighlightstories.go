package igstory

// Get highlight stories of a specific user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const UrlUserHighlightStories = `https://i.instagram.com/api/v1/highlights/{{USERID}}/highlights_tray/`

// Used to decode JSON returned by Instagram story API.
type RawHighlightsTray struct {
	Trays []HighlightTray `json:"tray"`
}

// Used to decode JSON returned by Instagram story API.
type HighlightTray struct {
	Id              string     `json:"id"`
	LatestReelMedia int64      `json:"latest_reel_media"`
	User            TrayUser   `json:"user"`
	Items           []TrayItem `json:"items"`
}

func GetUserHighlightStoriesTray(id int64) (trays []HighlightTray, err error) {
	url := strings.Replace(
		UrlUserHighlightStories,
		"{{USERID}}",
		strconv.FormatInt(id, 10),
		1)
	req, err := http.NewRequest("GET", url, nil)
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
	t := RawHighlightsTray{}
	if err = dec.Decode(&t); err != nil {
		return
	}
	trays = t.Trays

	return
}
