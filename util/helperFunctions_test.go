package util

import (
	"testing"
)

func TestSearchByName(t *testing.T) {
	channel := Channel{"en", "TestName", "testname", false, "Test Title", "Test Id", "Test Game", "Test Game Id", "Test Thumbnail", "Test Start Time"}
	var data []struct{ Channel }
	data = append(data, struct{ Channel }{channel})
	twitchChannels := TwitchChannels{data}
	_, err := SearchByName("testname", twitchChannels)
	if err != nil {
		t.Errorf("Could not find testname among the channels")
	}
}

func TestGetUserID(t *testing.T) {

}
