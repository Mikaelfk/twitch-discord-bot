package db

import (
	"context" // State handling across API boundaries; part of native GoLang API
	"log"

	"cloud.google.com/go/firestore"   // Firestore-specific support
	firebase "firebase.google.com/go" // Generic firebase support
	"google.golang.org/api/option"
)

// Firebase context and client used by Firestore functions throughout the program.
var ctx context.Context
var client *firestore.Client

// Collection name in Firestore
var collection = "Subscriptions"

type docFields struct {
	Channel_ids []string `firestore:"channel_ids,omitempty"`
}

// Gets all the channel ids by a streamer id
func GetChannelIdsByStreamerId(streamer_id string) []string {
	// Gets the document with the given streamer id
	doc, errNotFound := client.Collection(collection).Doc(streamer_id).Get(ctx)
	if errNotFound != nil {
		log.Println("Document not found")
		return nil
	}
	// Stores the data in a custom struct and returns the string slice with the ids
	var docData docFields
	doc.DataTo(&docData)
	return docData.Channel_ids
}

// Takes the streamer id and discord channel id as parameters and adds a subscription to the firestore
func AddSubscription(streamer_id string, channel_id string) error {
	// Tries to get the document with a matching streamer_id, if not found, adds a new document
	_, errNotFound := client.Collection(collection).Doc(streamer_id).Get(ctx)
	if errNotFound != nil {
		_, err := client.Collection(collection).Doc(streamer_id).Set(ctx, map[string]interface{}{
			"channel_ids": []interface{}{channel_id},
		})
		if err != nil {
			log.Println("Could not create new document")
			return err
		}
		return nil
	}
	// If the document exists, adds the channel id to the array in the document
	_, err := client.Collection(collection).Doc(streamer_id).Update(ctx, []firestore.Update{
		{
			Path:  "channel_ids",
			Value: firestore.ArrayUnion(channel_id),
		},
	})
	if err != nil {
		log.Println("Could not add to array")
		return err
	}
	return nil
}

// Deletes a subscription from the firestore
func DeleteSubscription(streamer_id string, channel_id string) error {
	// Tries to get the document with a matching streamer_id, if not found, returns an error
	_, errNotFound := client.Collection(collection).Doc(streamer_id).Get(ctx)
	if errNotFound != nil {
		log.Println("Document not found")
		return errNotFound
	}

	// If the document exists, removes the channel id from the array
	_, err := client.Collection(collection).Doc(streamer_id).Update(ctx, []firestore.Update{
		{
			Path:  "channel_ids",
			Value: firestore.ArrayRemove(channel_id),
		},
	})
	if err != nil {
		log.Println("Failed to remove channel id")
		return err
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
