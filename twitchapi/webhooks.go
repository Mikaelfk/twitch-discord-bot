package twitchapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"twitch-discord-bot/constants"
	"twitch-discord-bot/util"
)

type createWebhook struct {
	Type      string `json:"type"`
	Version   int    `json:"version"`
	Condition struct {
		BradcasterUserId string `json:"broadcaster_user_id"`
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
	Challenge string `json:"challenge"`
}

func StartListener() {
	http.HandleFunc("/", requestHandler)
	log.Println("Subscription functionality enabled, now listening on port ", util.Config.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(util.Config.Port), nil))
}

func CreateSubscription(userId string, subType string) {
	log.Println("Trying to add a new webhook for user: " + userId + " :D")

	body := new(createWebhook)
	// defaults for each req
	body.Version = 1
	body.Transport.Method = "webhook"
	body.Transport.Callback = util.Config.CallbackUrl
	// secret should probably be unique for the individual subscriptions, but this is better than nothing
	body.Transport.Secret = util.Config.TwitchWebhooksSecret

	// req specific
	body.Type = subType
	body.Condition.BradcasterUserId = userId

	// JSON IT!
	json, err := json.Marshal(body)
	if err != nil {
		log.Println("Error marshalling json, webhook not registered")
		return
	}

	// do request to twitch to request webhook
	var response creationConfirmation
	err = util.HandleRequest(constants.UrlTwitchWebhooks, http.MethodPost, &response, json)
	if err != nil {
		log.Println("Error while doing request to Twitch API, webhook not registered")
		return
	}

	// hopefully get verification pending response
	if response.Data[0].Status == "webhook_callback_verification_pending" {
		log.Println("Webhook subscription request successful. Verifying...")
	} else if response.Data[0].Status != "" {
		log.Println("Twitch did not like the webhook subscription request :(")
	} else {
		log.Println("Something went wrong trying to requ webhook subscription")
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// try to get twitch message type header
		messageType := r.Header.Get("Twitch-Eventsub-Message-Type")

		// check if it has twitch message type header
		if messageType == "webhook_callback_verification" {
			log.Println("Webhook verification request received. Starting webhook verification...")

			// try to get signature header
			reqSignature := r.Header.Get("Twitch-Eventsub-Message-Signature")

			if reqSignature != "" {
				// first need to verify signature with HMAC-SHA256
				//  The message is the concatenation of the Twitch-Eventsub-Message-Id header,
				// the Twitch-Eventsub-Message-Timestamp header, and the raw bytes of the request body.

				// get body bytes
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Println("Unable to read request body :(")
					log.Println("Aborting verification...")
					return
				}

				if verifySignature(r.Header.Get("Twitch-Eventsub-Message-Id"), r.Header.Get("Twitch-Eventsub-Message-Timestamp"), string(body), reqSignature) {
					// hell yeah it's verification time
					log.Println("Signatures match! Responding to verification request :D")

					// get json from verification request
					var verificationReq creationVerification
					err = json.Unmarshal([]byte(body), &verificationReq)
					if err != nil {
						log.Println("Unable to parse request body from verification request")
						log.Println("Aborting verification...")
						return
					}

					// return challenge
					w.WriteHeader(200)
					w.Write([]byte(verificationReq.Challenge))

					// ayo we did it
					log.Println("Webhook subscription added! :D")
					return
				} else {
					log.Println("Signatures do not match :O")
					log.Println("Aborting verification...")
					return
				}

			} else {
				log.Println("No signature, aborting verification")
				return
			}

		} else if messageType == "notification" {
			log.Print("AYO GOT A WEBHOOK!!!!!!!!!!! 🥳")
		} else {
			// Got a POST with the twitch header, but the header message is unknown
			log.Println("Recieved POST from (supposedly) Twitch with an unknown message type: " + messageType)
		}
	}
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
