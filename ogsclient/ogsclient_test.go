package ogsclient_test

import (
	"net/http"
	"testing"

	"github.com/phss/ogs-notifier/ogsclient"
	"github.com/stretchr/testify/assert"
)

func TestNewClient_userDetails(t *testing.T) {
	httpClient := &http.Client{}
	ogsClient := ogsclient.NewClient(httpClient, "somewhere")

	_, err := ogsClient.Me()

	assert.NotNil(t, err)
}
