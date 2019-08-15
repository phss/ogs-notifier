package main

import (
	"flag"
	"fmt"

	"github.com/phss/ogs-notifier/pkg/oauth"
	"github.com/phss/ogs-notifier/pkg/ogsclient"
)

func main() {
	clientID := flag.String("clientId", "", "")
	clientSecret := flag.String("clientSecret", "", "")
	username := flag.String("username", "", "")
	password := flag.String("password", "", "")
	flag.Parse()

	oauthConfig := oauth.Config{
		TokenURL:     "https://online-go.com/oauth2/token/",
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
	}
	oauthClient, err := oauth.PasswordCredentialsClient(oauthConfig, *username, *password)
	if err != nil {
		panic(err)
	}

	ogsClient := ogsclient.NewClient(oauthClient.HTTPClient, "http://online-go.com/api/v1/")

	user, err := ogsClient.Me.User()
	if err != nil {
		panic(err)
	}
	games, err := ogsClient.Me.Games()
	if err != nil {
		panic(err)
	}

	fmt.Printf("My ranking is %s\n", user.DisplayRanking())
	fmt.Println("My current games are:")
	for _, game := range *games {
		if !game.HasEnded() {
			fmt.Printf("- %s: %s vs %s\n", game.Name, game.Players.Black.Username, game.Players.White.Username)
		}
	}
}
