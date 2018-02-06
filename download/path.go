package igstorydl

import (
	"net/url"
	"path"
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

func BuildOutputPath(username, url string, timestamp int64) string {
	dirname := path.Join("stories", username)
	CreateDirIfNotExist(dirname)
	filename := path.Base(StripQueryString(url))
	ext := path.Ext(filename)
	t := time.Unix(timestamp, 0)
	ts := t.Format(time.RFC3339)
	p := path.Join(dirname, username+"-"+ts+ext)
	return p
}
