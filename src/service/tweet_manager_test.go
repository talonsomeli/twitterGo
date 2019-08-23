package service_test

import (
	"github.com/talonsomeli/src/domain"
	"github.com/talonsomeli/src/service"
	"strings"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)
	//tm.AddUser("pepe", "pepe@gmail.com", "pips", "3t5315g31hegds314")
	//service.AddUser("grupoesfera", "grupoesfera@grupoesfera.com", "grupoesfera", "536456462gdghf")

	user := "grupoesfera"
	text := "This is my first tweet"

	tweetActual := domain.NewTextTweet(user, text)
	// Operation
	id, _ := tm.PublishTweet(tweetActual)

	// Validation
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
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)

	var user string
	text := "This is my first tweet"

	tweetActual := domain.NewTextTweet(user, text)
	// Operation
	var err error

	_, err = tm.PublishTweet(tweetActual)

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
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)

	user := "pepe"
	var text string
	tweetActual := domain.NewTextTweet(user, text)
	// Operation
	var err error
	_, err = tm.PublishTweet(tweetActual)

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

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)

	user := "pepe"
	text := "This is my first tweet and second and third and fourth and fifth and sixth and seventh" +
		"and eight and ninth and tenth and eleventh and twelfth and thirteenth and fourteenth hahahahahahahahahahhahahahahahahahaha" +
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa dieeeeeeeeeeeeeeeeee"

	tweetActual := domain.NewTextTweet(user, text)
	// Operation
	var err error
	_, err = tm.PublishTweet(tweetActual)

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
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)

	user1, user2 := "pepe", "pepa"
	text1, text2 := "hola soy pepe", "y yo soy pepa"

	tweetActual := domain.NewTextTweet(user1, text1)
	segundoTweetActual := domain.NewTextTweet(user2, text2)
	// Operation
	id1, _ := tm.PublishTweet(tweetActual)
	id2, _ := tm.PublishTweet(segundoTweetActual)

	// Validation
	publishedTweets := tm.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[id1]
	secondPublishedTweet := publishedTweets[id2]

	if !isValidTweet(t, *firstPublishedTweet, id1, user1, text1) {
		return
	}
	// Same for secondPublishedTweet
	if !isValidTweet(t, *secondPublishedTweet, id2, user2, text2) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)

	var id int

	user := "grupoesfera"
	text := "This is my first tweet"
	tweetActual := domain.NewTextTweet(user, text)
	// Operation
	id, _ = tm.PublishTweet(tweetActual)

	// Validation
	publishedTweet := tm.GetTweetById(id)

	isValidTweet(t, *publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	primerTweet := domain.NewTextTweet(user, text)
	segundoTweet := domain.NewTextTweet(user, secondText)
	tercerTweet := domain.NewTextTweet(anotherUser, text)

	_, _ = tm.PublishTweet(primerTweet)
	_, _ = tm.PublishTweet(segundoTweet)
	_, _ = tm.PublishTweet(tercerTweet)
	// Operation
	count := tm.CountTweetsByUser(user)
	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tm := service.NewTweetManager(tweetWriter)
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	primerTweet := domain.NewTextTweet(user, text)
	segundoTweet := domain.NewTextTweet(user, secondText)
	tercerTweet := domain.NewTextTweet(anotherUser, text)

	// publish the 3 tweets
	_, _ = tm.PublishTweet(primerTweet)
	_, _ = tm.PublishTweet(segundoTweet)
	_, _ = tm.PublishTweet(tercerTweet)
	// Operation
	tweets := tm.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */ }
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
	isValidTweet(t, *firstPublishedTweet, 0, user, text)
	isValidTweet(t, *secondPublishedTweet, 1, user, text)
}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)
	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf("Tweet was not saved correctly")
	} else if savedTweet.GetId() != id {
		t.Errorf("Expected id: %d but was: %d", savedTweet.GetId(), id)
	}
}

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	// Create and publish a tweet
	var tweet domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	tweetManager.PublishTweet(tweet)
	// Operation

	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Errorf("Tweet was not saved correctly nor found")
	}
	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Tweet found is not the one searched for")
	}
}

func TestReadTweets(t *testing.T) {

	// Initialization
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data
	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)
	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	FileWriter := (tweetWriter).(*service.FileTweetWriter)
	savedTweets := FileWriter.Read()
	println(savedTweets)
	if savedTweets == "" {
		t.Errorf("Tweet was not saved correctly")
	} else if savedTweets != "@grupoesfera: This is my first tweet\n" {
		t.Errorf("Expected tweet was: @grupoesfera: This is my first tweet`, but was: `%v`", savedTweets)
	}
}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user string, text string) bool {
	return tweet.GetText() == text && tweet.GetUser() == user && tweet.GetId() == id
}

