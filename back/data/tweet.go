package data

type Tweets struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

var TweetList []*Tweets

func GetTweet() []*Tweets {
	return TweetList
}

func AddTweets(Tweet *Tweets) {
	TweetList = append(TweetList, Tweet)
}
