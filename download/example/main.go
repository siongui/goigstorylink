package main

import (
	"github.com/siongui/goigstorylink/download"
)

func main() {
	igstorydl.MonitorAndDownload("your user id", "your session id", "your csrftoken")
}
