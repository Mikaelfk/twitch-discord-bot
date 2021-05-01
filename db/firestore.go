package db

import (
	"context" // State handling across API boundaries; part of native GoLang API
	"errors"
	"fmt"
	"log"

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

//GetSubscription gets a subscription where the streamer and channel id matches
func GetSubscription(streamer string, channelId string) (string, string, error) {

	// Makes an iterator over all the documents in a collection
	iter := client.Collection(collection).Where("streamer_name", "==", streamer).Where("channel_id", "==", channelId).Documents(ctx)
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
		// This is somewhat unecessary as the iterator only iterates through documents where both instances occur
		if m["streamer_name"].(string) == streamer && m["channel_id"].(string) == channelId {
			return m["streamer_name"].(string), m["channel_id"].(string), nil
		}
	}

	// Returns empty strings and an error if no matches are found
	return "", "", errors.New("no matches")
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

// Gets the streamer name and discord channel id as parameters and adds a subscription to the firestore
func addSubscription(streamer_name string, channel_id string) error {

	_, _, err := client.Collection(collection).Add(ctx, map[string]interface{}{
		"streamer_name": streamer_name,
		"channel_id":    channel_id,
	})

	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

// Deletes a subscription from the firestore
func deleteSubscription(streamer_name string, channel_id string) error {
	// Get a collection where the streamer name and channel id are identical to the parameters
	iter := client.Collection(collection).Where("streamer_name", "==", streamer_name).Where("channel_id", "==", channel_id).Documents(ctx)

	// Iterate through the collection found
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		// Deletes the document
		_, err = client.Collection(collection).Doc(doc.Ref.ID).Delete(ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
			return err
		}
	}
	return nil
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
