package igstory

// Get all stories links of users that a user follows

type IGUser struct {
	Id       int
	Username string
	Stories  []IGStory
}

type IGStory struct {
	Timestamp int64
	Url       string
}

func GetUnreadStories() (users []IGUser, err error) {
	trays, err := GetReelsTray(config)
	if err != nil {
		return
	}

	for _, tray := range trays {
		user := IGUser{
			Id:       tray.Id,
			Username: tray.User.Username,
		}

		// One item represents one story
		for _, item := range tray.Items {
			// DO NOT use DeviceTimestamp. It's not reliable
			story := IGStory{
				Timestamp: item.TakenAt,
			}

			// check if the story is video or image
			if len(item.VideoVersions) > 0 {
				// the story is video
				story.Url = item.VideoVersions[0].Url
			} else {
				// the story is image
				story.Url = item.ImageVersions2.Candidates[0].Url
			}

			user.Stories = append(user.Stories, story)
		}

		users = append(users, user)
	}

	return
}

func SetUserId(s string) {
	config["ds_user_id"] = s
}

func SetSessionId(s string) {
	config["sessionid"] = s
}

func SetCsrfToken(s string) {
	config["csrftoken"] = s
}

// TODO
func GetAllStories() {
}
