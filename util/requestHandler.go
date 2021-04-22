package util

import (
	"encoding/json"
	"errors"
	"net/http"
)

// HandleRequest handles a request to the twitch API by the given URL and the specified rest-method 'method' ("POST", "GET")
// and decodes the response into the specified resType interface. This is so that you can decide what type-struct you want for each request.
func HandleRequest(URL string, method string, resType interface{}) error {


	if method != "GET" && method != "POST" {
		return errors.New("invalid method") //TODO: make 'method' constant?
	}


	client := &http.Client{}
	req, _ := http.NewRequest(method, URL, nil)
	req.Header.Set("client-id", Config.TwitchClientID)
	req.Header.Set("Authorization", "Bearer "+Config.TwitchAuthToken)
	res, _ := client.Do(req)

	json.NewDecoder(res.Body).Decode(&resType)

	return nil
}
