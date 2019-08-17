package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/phss/ogs-notifier/pkg/oauth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	Root.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticates the user user with OGS",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println(`Important: in order to authenticate with OGS, you need to provide
some information. These are the client ID, a client secret, your 
username and password. The client ID and secret can be obtained by
following the process in https://ogs.docs.apiary.io/#reference/authentication.
The username and password are, unfortunately, the same that you use
to login to OGS.

THIS CAN BE A SECURITY RISK!

This application will NOT store your username and password anywhere,
it will only use to initially authenticate and get a refresh token.
However, the refresh token WILL be stored under ~/.ogscli/config.json. 
Refresh tokens are short lived and allow the application to use 
the OGS API on behalf of yourself. Everytime you run a command,
the application will get a new refresh token, invalidating the 
previous one. It's important to note that this token could 
potentially be used to do bad things if compromised.`)

		fmt.Print("\nAre you comfortable with this? [yN]: ")
		answer, _ := reader.ReadString('\n')
		if strings.TrimSpace(answer) != "y" {
			os.Exit(0)
		}
		fmt.Print("\nClient ID: ")
		clientID, _ := reader.ReadString('\n')
		clientID = strings.TrimSpace(clientID)
		fmt.Print("Client Secret: ")
		clientSecret, _ := reader.ReadString('\n')
		clientSecret = strings.TrimSpace(clientSecret)
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		fmt.Print("Password: ")
		passwordBytes, _ := terminal.ReadPassword(int(syscall.Stdin))

		oauthConfig := oauth.Config{
			TokenURL:     "https://online-go.com/oauth2/token/",
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
		oauthClient, err := oauth.PasswordCredentialsClient(oauthConfig, username, string(passwordBytes))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.Set("clientId", clientID)
		viper.Set("clientSecret", clientSecret)
		viper.Set("refreshToken", oauthClient.Token.RefreshToken)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("\n\nDone!")
	},
}
