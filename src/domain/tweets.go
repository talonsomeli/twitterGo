package domain

import "time"

var id int = 0

type Tweet interface {
	GetId() int
	GetUser() string
	GetText() string
	String() string
	PrintableTweet() string
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	id int
}

 type ImageTweet struct {
 	TextTweet
 	Image string
 }

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func NewTextTweet(user string, text string) *TextTweet {
	date := time.Now()
	id++
	return &TextTweet{user, text, &date, id-1}
}

func NewImageTweet(user string, text string, URL string) *ImageTweet {
	date := time.Now()
	id++
	return &ImageTweet{TextTweet{user, text, &date, id-1},URL}
}

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet {
	date := time.Now()

	id++
	return &QuoteTweet{TextTweet{user, text, &date, id-1},quotedTweet}
}

//func NewImageTweet()

func (tw *TextTweet) GetId() int {
	return tw.id
}
func (tw *TextTweet) GetUser() string {
	return tw.User
}
func (tw *TextTweet) GetText() string {
	return tw.Text
}

func (tw *TextTweet) PrintableTweet() string {
	return "@" + tw.GetUser()+": " + tw.GetText()
}
func (tw *ImageTweet) PrintableTweet() string {
	return "@" + tw.GetUser() + ": " + tw.GetText() + " " + tw.Image
}
func (tw *QuoteTweet) PrintableTweet() string {
	return "@" + tw.GetUser() + ": " + tw.GetText() + ` "` + tw.QuotedTweet.PrintableTweet() + `"`
}

func (tw *TextTweet) String() string {
	return tw.PrintableTweet()
}