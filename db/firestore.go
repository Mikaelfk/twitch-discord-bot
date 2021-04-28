package db

import (
	"context" // State handling across API boundaries; part of native GoLang API
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"   // Firestore-specific support
	firebase "firebase.google.com/go" // Generic firebase support
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// Collection name in Firestore
var collection = "Subscriptions"

// Gets a subscription where the streamer and channel id matches
func GetSubscription(streamer string, channelId string) (string, string, error) {

	// Makes an iterator over all the documents in a collection
	iter := client.Collection(collection).Documents(ctx)
	for {
		// If the iterator is done, break out
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		// Returns an error if the iterator failed
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		// Stores document data in a map
		m := doc.Data()

		// Checks if both the streamer field and the channel id field matches the input
		if m["streamer_name"].(string) == streamer && m["channel_id"].(string) == channelId {
			return m["streamer_name"].(string), m["channel_id"].(string), nil
		}
	}

	// Returns empty strings and an error if no matches are found
	return "", "", errors.New("No matches")
}

// Get all subscriptions in a collection.
func GetSubscriptions() []string {

	// This array stores all the subscriptions
	var subscriptions []string

	iter := client.Collection(collection).Documents(ctx)
	for {
		// Iterates over every document found in the subscriptions collection.
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		m := doc.Data()
		// Appends every stream name entry to the subscriptions slice.
		subscriptions = append(subscriptions, m["streamer_name"].(string))
	}
	return subscriptions
}

// Reads a string from the body in plain-text and sends it to firestore to be registered as a message.
func AddSubscription(streamer_name string, channel_id string) {

	_, _, err := client.Collection("Subscriptions").Add(ctx, map[string]interface{}{
		"streamer_name": streamer_name,
		"channel_id":    channel_id,
	})

	if err != nil {
		fmt.Println(err)
		fmt.Println("\nsomething went wrong")

	}
}

// Deletes an individual or all messages from firestore
func deleteMessage(w http.ResponseWriter, r *http.Request) {

	// extract message ID from URL
	elem := strings.Split(r.URL.Path, "/")
	messageId := elem[2]

	// Delete individual message based on reference
	if len(messageId) != 0 {
		_, err := client.Collection(collection).Doc(messageId).Delete(ctx)
		if err != nil {
			http.Error(w, "Deletion of "+messageId+" failed.", http.StatusInternalServerError)
		}
		http.Error(w, "Deletion of "+messageId+" successful.", http.StatusNoContent)
	} else {
		// Delete all messages
		it := client.Collection(collection).Documents(ctx)
		for {
			item, err := it.Next()
			if err == iterator.Done {
				break
			}
			_, err = item.Ref.Delete(ctx)
			if err != nil {
				http.Error(w, "Deletion of item "+item.Ref.ID+" failed.", http.StatusInternalServerError)
			}
			http.Error(w, "Messages deleted.", http.StatusNoContent)
		}
	}
}

func InitDB() {
	// Firebase initialisation
	ctx = context.Background()

	// We use a service account, load credentials file that you downloaded from your project's settings menu.
	// Make sure this file is gitignored, it is the access token to the database.
	sa := option.WithCredentialsFile("service-account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	// defer client.Close()
}
