package util

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// JSON message structure used in tests.
type Message struct {
	Content        string `json:"content"`
	FurtherContent string `json:"furtherContent"`
}

// Starts a local test http server with fixed handler
func StartHTTPServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Populate header
		w.Header().Add("content-type", "application/json")
		// Write status code last (else previous headers will be reset)
		w.WriteHeader(http.StatusOK) // implicit value

		// Attach some content
		m := Message{"some content", "some more content"}
		content, err := json.Marshal(&m)
		if err != nil {
			log.Fatal("Error when marshaling: " + err.Error())
		}
		_, err = w.Write(content)
		if err != nil {
			log.Fatal("Error when writing content back: " + err.Error())
		}
	}))
}

// Tests the HandleRequest function with the use of local http server with a known response
func TestHandleRequest(t *testing.T) {
	var err error
	var message Message
	var url string

	// - tests network and handler, but abstracts from paths
	server := StartHTTPServer()
	// Ensure the server is closed at the end of the test
	defer server.Close()
	// Retrieve URL of test http server
	url = server.URL
	// Invoke with actual http default client
	err = HandleRequest(url, http.MethodGet, &message)

	// Check for general success
	if err != nil {
		t.Error("The request should have been successful, but received error instead: " + err.Error())
	}

	// Test content
	log.Println(message)
	if message.Content != "some content" {
		t.Errorf("Error in response content: %v", message.Content)
	}
}
