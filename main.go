package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)
go get -u 'github.com/openai/openai-go@v1.12.0'

var slackClient = slack.New(os.Getenv("SLACK_BOT_TOKEN"))
var APIkey = os.Getenv("OPENAI_API_KEY")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Slack sent a request!")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event, err := slackevents.ParseEvent(bytes, slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if event.Type == slackevents.URLVerification {
		challenge := event.Data.(*slackevents.ChallengeResponse)
		w.Write([]byte(challenge.Challenge))
		return
	}

	if event.Type == slackevents.CallbackEvent {
		innerEvent := event.InnerEvent

		if message, ok := innerEvent.Data.(*slackevents.MessageEvent); ok {
			fmt.Printf("[%s] %s: %s\n", message.Channel, message.User, message.Text)
			if message.BotID == "" {
				_, _, err := slackClient.PostMessage(
					message.Channel,
					slack.MsgOptionText(message.Text, false),
				)

				if err != nil {
					fmt.Println("Error posting message:", err)
				}
			}
		}
	}

	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/slack/events", handler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)

}
