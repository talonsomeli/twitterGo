package service

import (
	"fmt"
	"github.com/talonsomeli/src/domain"
	"strings"
)

type TweetManager struct {
	tweetsByUser map[string][]*domain.Tweet
	tweetsId map[int]*domain.Tweet
	writer TweetWriter
}

func NewTweetManager(tweetWriter TweetWriter) *TweetManager{
	var tm *TweetManager = new(TweetManager)
 	tm.tweetsByUser = make(map[string][]*domain.Tweet)
 	tm.tweetsId = make(map[int]*domain.Tweet, 0)
 	tm.writer = tweetWriter
 	return tm
}

/*func GetUsers() []*domain.User {
	return users
}
func AddUser(nombre string, mail string, nick string, clave string) {
	users = append(users, domain.NewUser(nombre, mail, nick, clave))
}*/

func (tm *TweetManager) PublishTweet(realTweet domain.Tweet) (int, error) {

	if len(realTweet.GetUser()) == 0 {
		return len(tm.tweetsId), fmt.Errorf("user is required")
	} else if len(realTweet.GetText()) == 0 {
		return len(tm.tweetsId), fmt.Errorf("text is required")
	} else if len(realTweet.GetText()) > 140 {
		return len(tm.tweetsId), fmt.Errorf("max tweet length is 140 characters")
	}

	tm.tweetsByUser[realTweet.GetUser()] = append(tm.tweetsByUser[realTweet.GetUser()], &realTweet)
	tm.tweetsId[realTweet.GetId()] = &realTweet
	tm.writer.Save(realTweet)
	return realTweet.GetId(), nil
}

func (tm *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {
	go func() {
		for _, tweet := range tm.tweetsId {
			if strings.Contains(((*tweet).(domain.Tweet)).GetText(), query) {
				searchResult <- *tweet
			}
		}
	}()
}

func (tm *TweetManager) GetTweets() map[int]*domain.Tweet {
	return tm.tweetsId
}

func (tm *TweetManager) GetTweetById(id int) *domain.Tweet {
	return tm.tweetsId[id]
}

func (tm *TweetManager) CountTweetsByUser(user string)int {

	return len(tm.tweetsByUser[user])
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {

	return tm.tweetsByUser[user]
}