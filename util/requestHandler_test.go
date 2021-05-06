package util

import (
	"net/http"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	var testResponse twitchUserSearch
	url := "https://api.twitch.tv/helix/users?login=ukhureaper"
	err := HandleRequest(url, http.MethodGet, &testResponse)
	if err != nil {
		t.Errorf("Request failed for Get request to url %v", url)
	}
}
