package service_test

import (
	"github.com/talonsomeli/src/domain"
	"github.com/talonsomeli/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization

	service.InitializeService()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ := service.PublishTweet(tweet)

	// Validation
	tweetActual := service.GetTweets()[0]
	if tweetActual.User != user ||
		tweetActual.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweetActual.User, tweetActual.Text)
	}
	if tweetActual.Date == nil {
		t.Error("Expected date can't be nil")
	}

	isValidTweet(t, tweetActual, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization

	service.InitializeService()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}


func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initialization

	service.InitializeService()
	var tweet *domain.Tweet

	user := "pepe"
	var text string
	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}
func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization

	service.InitializeService()
	var tweet *domain.Tweet

	user := "pepe"
	text := "This is my first tweet and second and third and fourth and fifth and sixth and seventh" +
		"and eight and ninth and tenth and eleventh and twelfth and thirteenth and fourteenth hahahahahahahahahahhahahahahahahahaha" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa dieeeeeeeeeeeeeeeeee"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error did not appear")
	}

	if err != nil && err.Error() != "max tweet length is 140 characters" {
		t.Error("Expected error is max tweet length is 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data

	user1, user2 := "pepe", "pepa"
	text1, text2 := "hola soy pepe", "y yo soy pepa"
	
	tweet = domain.NewTweet(user1, text1)
	secondTweet = domain.NewTweet(user2, text2)
	// Operation
	id1, _ := service.PublishTweet(tweet)
	id2, _ := service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, id1, user1, text1) {
		return
	}
	// Same for secondPublishedTweet
	if !isValidTweet(t, secondPublishedTweet, id2, user2, text2) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)
	// Operation
	count := service.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	service.InitializeService()
	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	// publish the 3 tweets
	_, _ = service.PublishTweet(tweet)
	_, _ = service.PublishTweet(secondTweet)
	_, _ = service.PublishTweet(thirdTweet)
	// Operation
	tweets := service.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */ }
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
	isValidTweet(t, firstPublishedTweet, 0, user, text)
	isValidTweet(t, secondPublishedTweet, 1, user, text)
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user string, text string) bool {
	return tweet.Text == text && tweet.User == user && domain.GetId(tweet) == id
}