package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

type QuoteResponse struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

var fedxLobby = "1124585337855950912"

var subscribedChannels = []string{"207372726221012993", "399577277903536138", fedxLobby}

func getRapidToken() string {
	token := os.Getenv("RAPID_API_TOKEN")
	if len(token) == 0 {
		return ""
	}
	return token
}

func getDurationUntil10Am() time.Duration {
	now := time.Now().UTC()

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(fmt.Sprintf("Failed to load location: %v", err))
	}
	now = now.In(loc)
	// Set the target time to 10 AM of the current day
	target := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())

	// If current time is after 10 AM, calculate for the next day's 9 AM
	if now.After(target) {
		target = target.Add(24 * time.Hour)
	}

	// Calculate the duration between now and the target time
	duration := target.Sub(now)

	return duration
}

func StartJob(s *discordgo.Session) {
	duration := getDurationUntil10Am()
	time.AfterFunc(duration, func() {
		QueryQuote(s)
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			QueryQuote(s)
		}
	})
}

func QueryQuote(s *discordgo.Session) {
	rapidAPIKey := getRapidToken()
	rapidAPIHost := "quotes-inspirational-quotes-motivational-quotes.p.rapidapi.com"

	url := "https://quotes-inspirational-quotes-motivational-quotes.p.rapidapi.com/quote"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	query := req.URL.Query()
	query.Add("token", "ipworld.info")
	req.URL.RawQuery = query.Encode()

	req.Header.Set("X-RapidAPI-Key", rapidAPIKey)
	req.Header.Set("X-RapidAPI-Host", rapidAPIHost)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var quoteResponse QuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		return
	}

	msg := fmt.Sprintf(" :bulb: Quote of the day :bulb:\n > %s", quoteResponse.Text)

	for _, c := range subscribedChannels {
		if c == fedxLobby {
			msg += "\n\n :tada: Happy Birthday, Justin! :tada:"
		}
		s.ChannelMessageSend(c, msg)
	}
}
