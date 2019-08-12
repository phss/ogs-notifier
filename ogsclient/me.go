package ogsclient

import (
	"github.com/dghubble/sling"
)

// MeService provides method for accessing resources under /me endpoint
// of the OGS API
type MeService struct {
	sling *sling.Sling
}

// User fetches the /me resource endpoint, which represents the authenticated user.
func (s *MeService) User() (*User, error) {
	user := new(User)
	apiError := new(ogsAPIError)
	_, err := s.sling.Get("me").Receive(user, apiError)
	return user, handleErrors(err, apiError)
}
