package igstory

// Get all stories links of users that a user follows

import (
	"fmt"
)

type IGUser struct {
	Id       int64
	Username string
	Stories  []IGStory
}

type IGStory struct {
	Timestamp int64
	Url       string
}

func itemsToStories(items []TrayItem) (stories []IGStory) {
	// One item represents one story
	for _, item := range items {
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

		stories = append(stories, story)
	}
	return
}

// Get stories of users with unread stories. The returned users will contains
// all users with unexpired stories, but only users with unread stories will
// have non-empty Stories field.
func GetUnreadStories() (users []IGUser, err error) {
	trays, err := GetReelsTray()
	if err != nil {
		return
	}

	for _, tray := range trays {
		user := IGUser{
			Id:       tray.Id,
			Username: tray.User.Username,
		}
		user.Stories = itemsToStories(tray.Items)

		users = append(users, user)
	}

	return
}

func fetchUserStories(user *IGUser, c chan int) {
	tray, err := GetUserStories(user.Id)
	if err != nil {
		// TODO: handle error here
		fmt.Println("In fetchUserStorie: fail to fetch " + user.Username)
		c <- 1
		return
	}
	user.Stories = itemsToStories(tray.Items)
	c <- 1
}

// Get all stories of users
func GetAllStories() (users []IGUser, err error) {
	users, err = GetUnreadStories()
	if err != nil {
		return
	}

	c := make(chan int)
	numOfEmptyStoryUser := 0
	for index, _ := range users {
		if len(users[index].Stories) == 0 {
			numOfEmptyStoryUser++
			go fetchUserStories(&users[index], c)
		}
	}

	// wait all goroutines to finish
	for i := 0; i < numOfEmptyStoryUser; i++ {
		<-c
	}

	return
}
