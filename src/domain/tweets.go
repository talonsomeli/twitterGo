package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
	id int
}

func NewTweet(user string, text string) *Tweet {
	date := time.Now()
	return &Tweet{user, text, &date, 0}
}

func NewTweetId(user string, text string, id int) *Tweet {
	date := time.Now()
	return &Tweet{user, text, &date, id}
}

func GetId(tweet *Tweet) int {
	return tweet.id
}