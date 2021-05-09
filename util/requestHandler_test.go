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
			"Test HandleRequest succeed",
			args{url, http.MethodGet, testResponse},
			false,
		},
		{
			"Test HandleRequest fail",
			args{"", http.MethodGet, testResponse},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if err := HandleRequest(tt.args.url, tt.args.method, tt.args.resType); (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
