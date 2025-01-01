package twitter

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dghubble/oauth1"
)

// SetupOAuth guides the user through Twitter OAuth authentication
func SetupOAuth(consumerKey, consumerSecret string) (string, string, error) {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		return "", "", err
	}

	authURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", err
	}

	fmt.Println("Please visit the following URL to authorize the app:")
	fmt.Println(authURL.String())
	fmt.Print("Enter the PIN provided by Twitter: ")

	reader := bufio.NewReader(os.Stdin)
	pin, _ := reader.ReadString('\n')

	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, pin)
	if err != nil {
		return "", "", err
	}

	return accessToken, accessSecret, nil
}
