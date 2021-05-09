package util

import (
	"reflect"
	"testing"
)

func TestSearchByName(t *testing.T) {
	type args struct {
		searchName string
		channels   TwitchChannels
	}

	channel := Channel{"en", "TestName", "testname", false, "Test Title", "Test Id", "Test Game", "Test Game Id", "Test Thumbnail", "Test Start Time"}
	var data []struct{ Channel }
	data = append(data, struct{ Channel }{channel})
	twitchChannels := TwitchChannels{data}

	tests := []struct {
		name    string
		args    args
		want    Channel
		wantErr bool
	}{
		{
			"Test 1",
			args{"testname", twitchChannels},
			channel,
			false,
		},
		{
			"Test 2",
			args{"random name", twitchChannels},
			Channel{},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
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
