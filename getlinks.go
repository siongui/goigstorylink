package igstory

// Get all stories links of users that a user follows

type IGUser struct {
	Id       int64
	Username string
	Stories  []IGStory
}

type IGStory struct {
	Timestamp int64
	Url       string
}

// Get stories of users with unread stories. The returned users will contains
// all users with unexpired stories, but only users with unread stories will
// have non-empty Stories field.
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

// Get all stories of users
func GetAllStories() (users []IGUser, err error) {
	users, err = GetUnreadStories()
	if err != nil {
		return
	}

	return
}
