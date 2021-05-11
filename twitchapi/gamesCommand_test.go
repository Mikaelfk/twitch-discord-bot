package twitchapi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)
/*
 JSON message structure used in tests game test.
*/
type Message struct {
	Data []data `json:"data"`
}
type data struct{
	ArtURL string `json:"box_art_url"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

var mName = "League of Legends"
var mID = "1121"
var mArt = "ART"


/*
Starts a local test http server with fixed handler
*/
func startHttpServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Populate header
		w.Header().Add("content-type", "application/json")
		// Write status code last (else previous headers will be reset)
		w.WriteHeader(http.StatusOK) // implicit value


		var g []data
		c := append(g, data{mArt,mID,mName})
		m := Message{Data: c}
		content, err := json.Marshal(&m)
		if err != nil {
			log.Fatal("Error when marshalling: " + err.Error())
		}
		_, err = w.Write(content)
		if err != nil {
			log.Fatal("Error when writing content back: " + err.Error())
		}
	}))
}


func TestFindGames(t *testing.T) {
	var url string

	// - tests network and handler, but abstracts from paths
	server := startHttpServer()
	// Ensure the server is closed at the end of the test
	defer server.Close()
	// Retrieve URL of test http server
	url = server.URL
	// Invoke with actual http default client

	gamesData, err := FindGames("League of Legends", 1, url)

	// Check for general success
	if err != nil {
		t.Error("The request should have been successful, but received error instead: " + err.Error())
	}

	// Test content
	log.Println(gamesData)

	if len(gamesData.Data)!=1{
		t.Errorf("Error in response content: %v", gamesData.Data)
	}

	if gamesData.Data[0].ID != mID {
		t.Errorf("Error in response content: %v", gamesData.Data[0].ID)
	}
	if gamesData.Data[0].Name != mName {
		t.Errorf("Error in response content: %v", gamesData.Data[0].Name)
	}
	if gamesData.Data[0].ArtURL != mArt {
		t.Errorf("Error in response content: %v", gamesData.Data[0].ArtURL)
	}

}

