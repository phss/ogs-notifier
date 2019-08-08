package ogsclient

import "errors"

type ogsAPIError struct {
	Detail string `json:"detail"`
}

func handleErrors(err error, apiError *ogsAPIError) error {
	if err != nil {
		return err
	}

	if apiError.Detail != "" {
		return errors.New(apiError.Detail)
	}

	return nil
}
