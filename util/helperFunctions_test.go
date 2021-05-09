package util

import (
	"reflect"
	"testing"
)

func TestSearchByName(t *testing.T) {
	channel := Channel{"en", "TestName", "testname", false, "Test Title", "Test Id", "Test Game", "Test Game Id", "Test Thumbnail", "Test Start Time"}
	var data []struct{ Channel }
	data = append(data, struct{ Channel }{channel})
	twitchChannels := TwitchChannels{data}

	type args struct {
		searchName string
		channels   TwitchChannels
	}

	tests := []struct {
		name    string
		args    args
		want    Channel
		wantErr bool
	}{
		{
			"Test Search By Name succeed",
			args{"testname", twitchChannels},
			channel,
			false,
		},
		{
			"Test Search By Name fail",
			args{"random name", twitchChannels},
			Channel{},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchByName(tt.args.searchName, tt.args.channels)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserID(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		t.Errorf("Could not load configuration, %v", err)
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Test GetUserID succeed 1",
			args{"twitch"},
			"12826",
			false,
		},
		{
			"Test GetUserID succeed 2",
			args{"ukhureaper"},
			"45432109",
			false,
		},
		{
			"Test GetUserID fail",
			args{""},
			"",
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserID(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
