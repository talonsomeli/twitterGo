package main

import (
	"github.com/abiosoft/ishell"
	"github.com/talonsomeli/src/domain"
	"github.com/talonsomeli/src/service"
	"strconv"
)


func main() {

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTextTweet(user, text)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes a tweet with an image",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			c.Print("Type the url of your image: ")

			url := c.ReadLine()

			tweet := domain.NewImageTweet(user, text, url)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishQuoteTweet",
		Help: "Publishes a tweet with a quote",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")

			user := c.ReadLine()

			c.Print("Type your tweet: ")

			text := c.ReadLine()

			c.Print("Type the id of the tweet you want to quote: ")

			id, _ := strconv.Atoi(c.ReadLine())

			quoteTweet := tweetManager.GetTweetById(id)

			tweet := domain.NewQuoteTweet(user, text, *quoteTweet)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows the last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := *tweetManager.GetTweets()[len(tweetManager.GetTweets()) - 1]

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetWriter.Read()

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet with the provided id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the id: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet := tweetManager.GetTweetById(id)

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts the tweets published by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the user: ")

			user := c.ReadLine()

			count := tweetManager.CountTweetsByUser(user)

			c.Println(count)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows the tweets published by the user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the user: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			c.Println(tweets)

			return
		},
	})

	shell.Run()

}