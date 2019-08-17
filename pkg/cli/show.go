package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/phss/ogs-notifier/pkg/oauth"
	"github.com/phss/ogs-notifier/pkg/ogsclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user's info",
	Run: func(cmd *cobra.Command, args []string) {
		clientID := viper.GetString("clientId")
		clientSecret := viper.GetString("clientSecret")
		refreshToken := viper.GetString("refreshToken")

		oauthConfig := oauth.Config{
			TokenURL:     "https://online-go.com/oauth2/token/",
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
		oauthClient, err := oauth.RefreshTokenClient(oauthConfig, refreshToken)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.Set("refreshToken", oauthClient.Token.RefreshToken)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ogsClient := ogsclient.NewClient(oauthClient.HTTPClient, "http://online-go.com/api/v1/")

		user, err := ogsClient.Me.User()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		games, err := ogsClient.Me.Games()
		if err != nil {
			panic(err)
		}

		fmt.Printf("My ranking is %s\n", user.DisplayRanking())
		fmt.Println("My current games are:")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "NAME\tBLACK\tWHITE\tLINK\t")
		for _, game := range *games {
			if !game.HasEnded() {
				fmt.Fprintf(w, "%s\t%s\t%s\thttps://online-go.com/game/%d\n", game.Name, game.Players.Black.Username, game.Players.White.Username, game.ID)
			}
		}
		w.Flush()
	},
}
