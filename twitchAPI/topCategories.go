package twitchAPI

import (
	"errors"
	"log"
	"twitch-discord-bot/util"
)

// A struct that stores all the category names
type TopCategoriesResult struct {
	Data []struct {
		Name string `json:"name"`
	} `json:"data"`
}

// Gets the current top 13 categories on twitch
func GetTopCategories() ([]string, error) {

	// Stores all the categories
	var categories []string
	var topCategories TopCategoriesResult

	// Calls a GET request to top games from the twitch API
	err := util.HandleRequest("https://api.twitch.tv/helix/games/top", "GET", &topCategories)

	// Returns an error if it couldn't parse the request into the struct
	if err != nil {
		log.Fatal(err)
	}

	// Returns an error if there are no categories
	if topCategories.Data == nil {
		return nil, errors.New("no categories found")
	}

	// Iterates through each category in the topCategories array
	for i, category := range topCategories.Data {

		categories = append(categories, category.Name)

		// Stops iterating when the index reaches 12, this assures that it only returns the top 13 categories
		if i == 12 {
			break
		}
	}

	return categories, nil
}
