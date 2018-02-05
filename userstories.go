package igstory

// Get all stories of a specific user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const UrlUserStories = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/reel_media/`

func GetUserStories(id string, cfg map[string]string) (tray Tray, err error) {
	url := strings.Replace(UrlUserStories, "{{USERID}}", id, 1)
	req, err := http.NewRequest("GET", url, nil)
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
	err = dec.Decode(&tray)
	return
}
