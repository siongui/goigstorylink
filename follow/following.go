package igstory

// Get all following users

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const UrlFollowing = `https://i.instagram.com/api/v1/friendships/{{USERID}}/following/`

type RawFollowing struct {
	Users    []FollowUser `json:"users"`
	BigList  bool         `json:"big_list"`
	PageSize int64        `json:"page_size"`
	Status   string       `json:"status"`
}

type FollowUser struct {
	Pk       int64  `json:"pk"`
	Username string `json:"username"`
}

func GetFollowing(ds_user_id, sessionid, csrftoken string) (rf RawFollowing, err error) {
	url := strings.Replace(UrlFollowing, "{{USERID}}", ds_user_id, 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: ds_user_id})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: sessionid})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: csrftoken})

	req.Header.Set("User-Agent", "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+")

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
	err = dec.Decode(&rf)
	return
}
