package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

var (
	// ChannelSecret is LINE Bot Channel Secret.
	ChannelSecret = os.Getenv("CHANNEL_SECRET")
	// ChannelAccessToken is LINE Bot Channel Access Token.
	ChannelAccessToken = os.Getenv("CHANNEL_ACCESS_TOKEN")
	// UserID is LINE Bot Your user ID.
	UserID = os.Getenv("USER_ID")
)

func main() {
	log.Print("===== yaruki application start =====")
	bot, err := linebot.New(
		ChannelSecret,
		ChannelAccessToken,
	)
	if err != nil {
		log.Print(err)
		return
	}
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if _, err = bot.PushMessage(UserID, linebot.NewTextMessage("今日もpushしてえらい！(*´>ω<))ω｀●)")).Do(); err != nil {
			log.Print(err)
			// HTTP 500 (Internal Server Error)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// HTTP 200 (OK)
		w.WriteHeader(http.StatusOK)
	})
	log.Print(http.ListenAndServe(":8080", nil))
}