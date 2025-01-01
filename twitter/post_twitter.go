package twitter

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

func PostTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, message string) {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	req, _ := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer([]byte(`{"text":"`+message+`"}`)))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Failed to send tweet:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		fmt.Println("Successfully posted tweet!")
	} else {
		fmt.Println("Failed to post tweet. Status:", resp.Status)
	}
}
