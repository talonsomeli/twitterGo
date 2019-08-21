package service

import (
	"fmt"
	"github.com/talonsomeli/src/domain"
)

var tweetsUser map[int]*domain.Tweet
var tweetsId []*domain.Tweet
func InitializeService() {
 tweetsUser = make(map[int]*domain.Tweet)
 tweetsId = make([]*domain.Tweet, 0)
}

func PublishTweet(tweetToPublish *domain.Tweet) (int, error) {

	if len(tweetToPublish.User) == 0 {
		return len(tweetsId), fmt.Errorf("user is required")
	} else if len(tweetToPublish.Text) == 0 {
		return len(tweetsId), fmt.Errorf("text is required")
	} else if len(tweetToPublish.Text) > 140 {
		return len(tweetsId), fmt.Errorf("max tweet length is 140 characters")
	}
	tweetId := len(tweetsId)
	realTweet := domain.NewTweetId(tweetToPublish.User, tweetToPublish.Text, tweetId)
	tweetsUser[tweetId] = realTweet
	tweetsId = append(tweetsId, realTweet)
	return tweetId, nil
}

func GetTweets() []*domain.Tweet {
	return tweetsId
}

func GetTweetById(id int) *domain.Tweet {
	return tweetsId[id]
}

func CountTweetsByUser(user string)int {

	cant := 0
	for _, valor := range tweetsUser {
		if valor.User == user {cant++}
	}
	return cant
}

func GetTweetsByUser(user string) []*domain.Tweet {

	var tweets []*domain.Tweet
	for _, valor := range tweetsUser {
		if valor.User == user {
			tweets = append(tweets, valor)
		}
	}
	return tweets
}