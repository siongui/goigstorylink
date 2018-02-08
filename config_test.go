package igstory

import (
	"testing"
)

func TestSetUserId(t *testing.T) {
	SetUserId("11111111")
	if config["ds_user_id"] != "11111111" {
		t.Error("func SetUserId fail")
	}
}

func TestSetSessionId(t *testing.T) {
	SetSessionId("11111111")
	if config["sessionid"] != "11111111" {
		t.Error("func SetSessionId fail")
	}
}

func TestSetCsrfToken(t *testing.T) {
	SetCsrfToken("11111111")
	if config["csrftoken"] != "11111111" {
		t.Error("func SetCsrfToken fail")
	}
}
