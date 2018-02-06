package igstorydl

import (
	"net/url"
	"path"
	"strconv"
	"time"
)

func StripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}

func FormatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func BuildOutputFilePath(username, url string, timestamp int64) string {
	dirname := path.Join("stories", username)
	CreateDirIfNotExist(dirname)
	filename := path.Base(StripQueryString(url))
	ext := path.Ext(filename)
	ts := FormatTimestamp(timestamp)
	p := path.Join(dirname, username+"-"+ts+"-"+strconv.FormatInt(timestamp, 10)+ext)
	return p
}
