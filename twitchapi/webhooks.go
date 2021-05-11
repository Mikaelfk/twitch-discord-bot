package twitchapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/db"
	"twitch-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

// bot session for sending messages and such
var discordSession *discordgo.Session

var creationMap map[string]func(bool)

type webhook struct {
	Type      string `json:"type"`
	Version   string `json:"version"`
	Condition struct {
		BradcasterUserID string `json:"broadcaster_user_id"`
	} `json:"condition"`
	Transport struct {
		Method   string `json:"method"`
		Callback string `json:"callback"`
		Secret   string `json:"secret"`
	} `json:"transport"`
}

type creationConfirmation struct {
	Data []struct {
		Status string `json:"status"`
	} `json:"data"`
}

type creationVerification struct {
	Challenge    string  `json:"challenge"`
	Subscription webhook `json:"subscription"`
}

type receivedWebook struct {
	Subscription webhook `json:"subscription"`
}

// StartListener starts a http server for registering incoming webhook stuff
func StartListener(session *discordgo.Session) {
	creationMap = make(map[string]func(bool))
	discordSession = session
	http.HandleFunc("/", requestHandler)
	log.Println("Subscription functionality enabled, now listening on port", util.Config.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(util.Config.Port), nil))
}

func CreateSubscription(userId string, subType string, callBackFunc func(bool)) {
	creationMap[userId] = callBackFunc

	log.Println("Trying to add a new webhook for user: " + userId + " :D")

	_, err := db.GetChannelIdsByStreamerID(userId)
	if err != nil {
		log.Println("Webhook for this streamer does not exist, trying to create...")
	} else {
		creationMap[userId](true)
		delete(creationMap, userId)
		return
	}

	body := new(webhook)
	// defaults for each req
	body.Version = "1"
	body.Transport.Method = "webhook"
	body.Transport.Callback = util.Config.CallbackURL
	// secret should probably be unique for the individual subscriptions, but this is better than nothing
	body.Transport.Secret = util.Config.TwitchWebhooksSecret

	// req specific
	body.Type = subType
	body.Condition.BradcasterUserID = userId

	// JSON IT!
	json, err := json.Marshal(body)
	if err != nil {
		log.Println("Error marshaling json, webhook not registered")
		creationMap[userId](false)
		delete(creationMap, userId)
	}

	// do request to twitch to request webhook
	var response creationConfirmation
	err = util.HandleRequest(constants.ULTwitchWebhooks, http.MethodPost, &response, json)
	if err != nil {
		log.Println("Error while doing request to Twitch API, webhook not registered")
		creationMap[userId](false)
		delete(creationMap, userId)
	}

	switch status := response.Data[0].Status; {
	case status == "webhook_callback_verification_pending":
		log.Println("Webhook subscription request successful. Verifying...")

		// if it's still in map after 60 seconds, something failed and it's time to abort
		time.AfterFunc(1*time.Minute, func() {
			if _, ok := creationMap[userId]; ok {
				creationMap[userId](false)
				delete(creationMap, userId)
			}
		})
	case status != "webhook_callback_verification_pending":
		log.Println("Twitch did not like the webhook subscription request :(")
		creationMap[userId](false)
		delete(creationMap, userId)
	default:
		log.Println("Something went wrong trying to requ webhook subscription")
		creationMap[userId](false)
		delete(creationMap, userId)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		messageType := r.Header.Get("Twitch-Eventsub-Message-Type")

		// check if it has twitch message type header
		switch messageType {
		case "webhook_callback_verification":
			log.Println("Webhook verification request received. Starting webhook verification...")

			// try to get signature header
			reqSignature := r.Header.Get("Twitch-Eventsub-Message-Signature")
			if reqSignature != "" {
				// get body bytes
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Unable to read request body :(")
					log.Println("Aborting verification...")
					return
				}

				if verifySignature(r.Header.Get("Twitch-Eventsub-Message-Id"), r.Header.Get("Twitch-Eventsub-Message-Timestamp"), string(body), reqSignature) {
					log.Println("Signatures match! Responding to verification request :D")

					respondToVerification(w, body)
				} else {
					log.Println("Signatures do not match :O")
					log.Println("Aborting verification...")
					return
				}
			} else {
				log.Println("No signature, aborting verification")
				return
			}
		case "notification":
			log.Print("Received a webhook, verifying...")

			// try to get signature header
			reqSignature := r.Header.Get("Twitch-Eventsub-Message-Signature")

			if reqSignature != "" {
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Unable to read webhook body :(")
					log.Println("Aborting received webhook verification...")
					return
				}

				if verifySignature(r.Header.Get("Twitch-Eventsub-Message-Id"), r.Header.Get("Twitch-Eventsub-Message-Timestamp"), string(body), reqSignature) {
					log.Println("Webhook verified!")

					// confirm to twitch that webhook has been received and verified
					w.WriteHeader(200)

					err := handleWebhook(w, body)
					if err != nil {
						log.Println("Error occurred, no notifications sent :(")
					}
				}
			}
		default:
			// Got a POST with the twitch header, but the header message is unknown
			log.Println("Received POST from (supposedly) Twitch with an unknown message type: " + messageType)
		}
	}
}

func respondToVerification(w http.ResponseWriter, body []byte) {
	// get json from verification request
	var verificationReq creationVerification
	err := json.Unmarshal([]byte(body), &verificationReq)
	if err != nil {
		log.Print(err)
		log.Println("Unable to parse request body from verification request")
		log.Println("Aborting verification...")
		return
	}

	// return challenge
	w.WriteHeader(200)
	w.Write([]byte(verificationReq.Challenge))

	// ayo we did it
	log.Println("Webhook subscription added! :D")
	if val, ok := creationMap[verificationReq.Subscription.Condition.BradcasterUserID]; ok {
		val(true)
		delete(creationMap, verificationReq.Subscription.Condition.BradcasterUserID)
	}
}

// handleWebhook will check contents of the webhook, and send notifications to subscribed channels if there are any
func handleWebhook(w http.ResponseWriter, body []byte) error {
	// get json from webhook request
	var recWebhook receivedWebook
	err := json.Unmarshal([]byte(body), &recWebhook)
	if err != nil {
		log.Println("Unable to parse request body from webhook")
		log.Println("Aborting trying to send notification")
		return err
	}

	// try to get subscribed channels from firestore
	channels, err := db.GetChannelIdsByStreamerID(recWebhook.Subscription.Condition.BradcasterUserID)
	if err != nil {
		log.Println("Unable to get channels for streamer, no notications sent")
		log.Println("If you see this often, there might be some discrepency between subscriptions and firebase")
		return err
	} else if len(channels) == 0 {
		log.Println("No channels registered for streamer, no notications sent")
		log.Println("If you see this often, there might be some discrepency between subscriptions and firebase")
		return errors.New("no channels registered for streamer")
	}

	// check if streamer is live, so no spam if they go online then quickly offline again
	stream, err := util.GetStreamDetails(recWebhook.Subscription.Condition.BradcasterUserID)
	if len(stream.Data) == 0 || err != nil {
		log.Println("Unable to get stream data, maybe they didn't actually go live? :O")
		return errors.New("unable to get stream data")
	}

	// send notifications to channels
	for i := range channels {
		// create cool discord embed
		var em discordgo.MessageEmbed
		em.Type = discordgo.EmbedType("rich")
		em.URL = constants.URLTwitchStream + stream.Data[0].UserLogin
		em.Color = 1

		if stream.Data[0].GameName == "" {
			em.Fields = []*discordgo.MessageEmbedField{{Name: "Game", Value: "No game", Inline: true}}
		} else {
			em.Fields = []*discordgo.MessageEmbedField{{Name: "Game", Value: stream.Data[0].GameName, Inline: true}}
		}

		if stream.Data[0].Title == "" {
			em.Title = "Stream"
		} else {
			em.Title = stream.Data[0].Title
		}

		// set dimensions for thumbnail
		thumbnailURL := stream.Data[0].ThumbnailUrl
		thumbnailURL = strings.ReplaceAll(thumbnailURL, "{width}", "640")
		thumbnailURL = strings.ReplaceAll(thumbnailURL, "{height}", "480")

		em.Image = &discordgo.MessageEmbedImage{
			URL: thumbnailURL,
		}

		// try to send noticiations
		_, err := discordSession.ChannelMessageSend(channels[i], stream.Data[0].UserName+" just went live!")
		if err != nil {
			log.Println("Unable to send notification to channel " + channels[i])
		}
		_, err = discordSession.ChannelMessageSendEmbed(channels[i], &em)
		if err != nil {
			log.Println("Unable to send embed to channel " + channels[i])
		}
	}
	return nil
}

func verifySignature(messageIDHeader string, timestampMessage string, body string, reqSignature string) bool {
	// create message to hash
	hmacMessage := []byte(fmt.Sprintf("%s%s%s", messageIDHeader, timestampMessage, body))

	// create hash with secret
	hmac := hmac.New(sha256.New, []byte(util.Config.TwitchWebhooksSecret))
	hmac.Write(hmacMessage)
	signature := fmt.Sprintf("sha256=%s", hex.EncodeToString(hmac.Sum(nil)))
	return signature == reqSignature
}
