package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Tweet struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}

type TweetHandler struct {
	mu     sync.Mutex
	tweets []Tweet
	log    *log.Logger
	nextID int
}

func NewTweet(logger *log.Logger) *TweetHandler {
	return &TweetHandler{
		tweets: []Tweet{},
		log:    logger,
		nextID: 1,
	}
}

func (th *TweetHandler) AddTweets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var tweet Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	th.mu.Lock()
	tweet.ID = th.nextID
	th.nextID++
	th.tweets = append(th.tweets, tweet)
	th.mu.Unlock()

	w.WriteHeader(http.StatusOK)
}

func (th *TweetHandler) GetTweet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	th.mu.Lock()
	defer th.mu.Unlock()

	// Render tweets as HTML
	var htmlContent string
	for _, tweet := range th.tweets {
		htmlContent += `<div class="p-4 bg-white mb-4 rounded shadow">
            <p>` + tweet.Content + `</p>
            <div class="mt-2">
                <button hx-post="/like-tweet?id=` + strconv.Itoa(tweet.ID) + `" class="px-4 py-2 bg-gray-200 rounded">like</button>
                <span>` + strconv.Itoa(tweet.Likes) + ` likes</span>
            </div>
        </div>`
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}
