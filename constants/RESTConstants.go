package constants

const (
	UrlTwitchChannelId   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	UrlTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	UrlTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
<<<<<<< HEAD
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
	UrlTwitchTopGames    = "https://api.twitch.tv/helix/games/top"
>>>>>>> Refactored the GetTopCategories function so that it uses a rest constant, instead of a hard coded string.
)
