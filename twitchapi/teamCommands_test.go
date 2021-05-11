package twitchapi

import (
	"reflect"
	"testing"
	"twitch-discord-bot/util"
)

func TestGetLiveTeamMembers(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"test tsm false",
			args{"tsm"},
			[]string{"nothing"},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLiveTeamMembers(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveTeamMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveTeamMembers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllTeamMembers(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"test tsm members true",
			args{"tsm"},
			[]string{},
			true,
		},
		{
			"test tsm members false",
			args{"tsm"},
			[]string{"nothing"},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllTeamMembers(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTeamMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTeamMembers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeamExist(t *testing.T) {
	util.LoadConfig("../")
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"tsm test 1 lol",
			args{"tsm"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := TeamExist(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("TeamExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TeamExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
