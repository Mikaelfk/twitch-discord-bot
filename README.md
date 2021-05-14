# Twitch-Discord Bot

Invite link to a server where you can try it out: [https://discord.gg/WGbmR6TVxE](https://discord.gg/WGbmR6TVxE)

## Description

The original project plan was to create a Discord Bot that integrates with Twitch. It would function as a bridge between Discord and Twitch for many Discord users, especially Discord server owners. The owners could use it to share news and notifications related to Twitch with hundreds of people simultaneously, making Twitch more accessible for a larger audience. Some of the first ideas of functionality were the ability to get info about streams (live viewers, title, Etc.), information about a particular channel, webhook notifications for when a specific stream goes live, and the ability to notify of drop events. 

According to the original plan this project would use the following technologies:
- Twitch API
- Discord API
- Firebase
- OpenStack / Heroku

## Final Product

The final project uses all of the originally planned technologies apart from OpenStack / Heroku, as we landed on using Docker for deployment instead. Heroku would have been a good solution as it would make it easy to aquire a HTTPS callback url, but since it gives a new port for each time the application runs, it would break the webhook notifications.

The bot has the following features represented by the following commands:
- **/channel [:login-name]** : `Gets the info registered by the Twitch API about a particular channel/streamer by their [:login-name].`
- **/follow-list [:username]** : `Returns a list of all registered Twitch users a [:username] follows.`
- **/games [:game-name] [?num-games]** : `Returns the full name, ID and thumbnail icon of a game by the specified [:game-name]. Users can also specify a  [?num-games] to show, but the default is one game as the response.`
- **/subscribe [:streamer-name]** : `Will set up going-live notifications in the current channel for a streamer by their login/username`
- **/unsubscribe [:streamer-name]** `Will remove the going-live notification for a streamer in the current channel`
- **/stream [?login-name] [?game-name] [?language] [?game-id]** : `will return the top 3 currently most viewed streams. Optionally, a user can specify a search for streams by login-name, game-name, language or game id. User can choose either some or all of the options.`
- **/top-categories** : `Will return the current top categories by their popularity on Twitch.`
- **/team [:team-name] [:is-live]** : `Will return all the members of a Twitch-team. User can decide if they want all or only live members.`

We, however, did not implement a drop event notification feature.

_**[?parameter]** represents an optional parameter, whereas **[:parameter]** represents a mandatory required parameter for the command_

## Config

### config.json

This is the main config file for the bot. Fill out the config template `config_template.json` and save it as `config.json` in the root directory.
- DiscordBotToken
    - This is the bot token acquired from Discord after registeing a bot.
- DiscordServerID (optional)
    - The ID of the Discord server to register commands in.
    - If left blank: All commands will be registered in all severs the bot is in. 
- TwitchClientID
    - This is the client ID acquired from Twitch after registering an application.
- TwitchAuthToken
    - This is the authentication token acquired from Twitch after registering an application
- TwitchWebhooksSecret
    - This is the secret used for verifying webhooks. Can be anything, as long as it's not easily guessable.
- EnableSubscriptionsFunctionality
    - This is an option to enable or disable subscription funcionality.
- CallbackURL
    - This is the callback url to use for webhooks. **The callback url must use HTTPS, as the Twitch API requires this!**
- Port
    - This is the port the bot is going to be listening for webhooks on.

### service-account.json

This is the authentication file for the firebase connection. It is exported from the firebase web-ui and must be saved as `service-account.json` in the root directoy. In the Firestore, a document named `Subscriptions` must also be created before use.

## Deployment

### Docker 

This app can be deployed with Docker. After getting `config.json` and `service-account.json` ready, you can use these two commands to run the bot:
- `docker build -t twitch-discord-bot .`
- `docker run -dp $PORT:$PORT twitch-discord-bot`
    - Replace $PORT with the same port used for webhook callbacks in the main config file

### Other

If you don't want to use Docker, it is possible to also simply run the bot by simply running `go run main.go`, or building it first with `go build`, followed by `./twitch-discord-bot`. Be aware that `config.json` and `service-account.json` still need to be present in the working directory.

It *should* run on Windows, Linux and MacOS :)

## What went well and what did not?

The teamwork and individual effort from everyone in this group have been excellent, resulting in quality work being done by each one. Despite not having a detailed time management plan, we had meetings regularly, which ensured that everyone was engaged and had something to do. The structure of the project work, however, ended up being perhaps a bit messy. We tried using issue-tracker and individual branches for each issue, but admittedly it could have been just a little bit better and cleaner. Moreover, the lack of a specific plan from the start made it so that it was hard to know what features to add. We only had a very vague idea, but we should perhaps have used more time to discuss the features we wanted to add before adding them. However, it turned out reasonably good, and it could have been further improved with more time.

## What was hard?

Some of the most challenging aspects of the project may have been webhooks, mainly the HTTPS callback url requirement of the Twitch API and the webhook verification process. Additionally, the authentication with the Twitch and Discord APIs was also tricky. However, we solved it reasonably fast as some in our group had prior experience with it. Even though it was challenging, we also learned a lot from it. 

## What we leaned?

We learned more about Firebase and testing in Golang. New things we learned were to use the Twitch and Discord APIs with their different authentication methods, and we learned about webhook invocation and verification. Deploying with Docker was also a new learning experience for us. 


## Total Hours of work
171 hours total between all the group members
