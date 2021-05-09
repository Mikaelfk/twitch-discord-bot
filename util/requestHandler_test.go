package util

import (
	"net/http"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	var testResponse twitchUserSearch
	url := "https://api.twitch.tv/helix/"

	type args struct {
		url     string
		method  string
		resType interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test Handle Request succeed",
			args{url, http.MethodGet, testResponse},
			false,
		},
		{
			"Test Handle Request fail",
			args{"", http.MethodGet, testResponse},
			true,
		},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			if err := HandleRequest(tests[i].args.url, tests[i].args.method, tests[i].args.resType); (err != nil) != tests[i].wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tests[i].wantErr)
			}
		})
	}
}
