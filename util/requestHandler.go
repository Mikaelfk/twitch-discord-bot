package util

import (
	"encoding/json"
	"errors"
	"net/http"
)

// HandleRequest handles a request to the twitch API by the given URL (NO SPACES ALLOWED) and the specified rest-method 'method' ("POST", "GET")
// and decodes the response into the specified resType interface. This is so that you can decide what type-struct you want for each request.
// Note: even if this method returns nil, it does not guarantee it decoded correctly into the specified resType. It just means it had no decoding errors.
// Therefore, if the contents of the resType is important, make sure to check for empty values.
//
// IMPORTANT: Do not include spaces in the URL as it will make this method return EOF error.
func HandleRequest(URL string, method string, resType interface{}) error {

	// only valid methods allowed. Can add more methods later if need be
	if method != http.MethodGet && method != http.MethodPost && method != http.MethodPatch && method != http.MethodDelete {
		return errors.New("invalid method")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return err
	}

	// set the headers so that the twitch API accepts the request
	req.Header.Set("client-id", Config.TwitchClientID)
	req.Header.Set("Authorization", "Bearer "+Config.TwitchAuthToken)
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&resType)

	// returns the error, which is nil if everything is ok. However, this doesn't mean that the resType actually has the desired data (additional checks needed).
	return err
}
