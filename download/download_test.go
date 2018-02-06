package igstorydl

import (
	"testing"
)

func ExampleMonitorAndDownload(t *testing.T) {
	MonitorAndDownload("your user id", "your session id", "your csrftoken")
}
