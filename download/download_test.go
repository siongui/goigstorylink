package igstorydl

import (
	"testing"
)

func TestMonitorAndDownload(t *testing.T) {
	MonitorAndDownload("your user id", "your session id", "your csrftoken")
}
