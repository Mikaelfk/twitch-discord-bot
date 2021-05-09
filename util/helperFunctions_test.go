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
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			got, err := SearchByName(tests[i].args.searchName, tests[i].args.channels)
			if (err != nil) != tests[i].wantErr {
				t.Errorf("SearchByName() error = %v, wantErr %v", err, tests[i].wantErr)
				return
			}
			if !reflect.DeepEqual(got, tests[i].want) {
				t.Errorf("SearchByName() = %v, want %v", got, tests[i].want)
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
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			got, err := GetUserID(tests[i].args.username)
			if (err != nil) != tests[i].wantErr {
				t.Errorf("GetUserID() error = %v, wantErr %v", err, tests[i].wantErr)
				return
			}
			if got != tests[i].want {
				t.Errorf("GetUserID() = %v, want %v", got, tests[i].want)
			}
		})
	}
}
