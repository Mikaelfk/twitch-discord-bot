package constants

<<<<<<< HEAD
const (
	UrlTwitchChannelId   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	UrlTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	UrlTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
	UrlTwitchStreamInfo  = "https://api.twitch.tv/helix/streams"
	UrlTwitchGames       = "https://api.twitch.tv/helix/search/categories?query=" // + [game_name]
	UrlTwitchTopGames    = "https://api.twitch.tv/helix/games/top"

	TwitchApiResolution     = "52x72.jpg"
	DiscordBotImgResolution = "200x150"

	ParaUserLogin = "user_login="
	ParaGameId    = "game_id="
	ParaLanguage  = "language="
)

// Discord Bot Error Messages
const (
	BotUnexpectedErrorMsg = "I'm vewy sorwy but somwthing wierd happened... >0<"
	BotNoResultsMsg       = "I'm so sowwy... I didn't find anything... (o.O)"
	BotNoActiveStreamsMsg = "Sowwy... There doesn't seem to be a stream like that active at the moment <0.o>"
	BotNoGames            = "...I'm afraid there are no games like that.. sowwy >.<"
=======
const (
	URL_TWITCH_CHANNEL_ID   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	URL_TWITCH_CHANNEL_NAME = "https://api.twitch.tv/helix/search/channels?query="   // + [name]
	URL_TWITCH_USER_NAME    = "https://api.twitch.tv/helix/users?login="
	URL_TWITCH_FOLLOWLIST   = "https://api.twitch.tv/helix/users/follows?from_id="
>>>>>>> Almost finished getting all the followed streamers, still WIP
)
