// Package constants provides often used values
package constants

const (
	URLTwitchChannelID   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	URLTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	URLTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
	URLTwitchStreamInfo  = "https://api.twitch.tv/helix/streams"
	URLTwitchGames       = "https://api.twitch.tv/helix/search/categories?query=" // + [game_name]
	URLTwitchTopGames    = "https://api.twitch.tv/helix/games/top"
	URLTwitchUserName    = "https://api.twitch.tv/helix/users?login="
	URLTwitchFollowlist  = "https://api.twitch.tv/helix/users/follows?from_id="

	TwitchApiResolution     = "52x72.jpg"
	DiscordBotImgResolution = "200x150"

	ParaUserLogin = "user_login="
	ParaGameId    = "game_id="
	ParaLanguage  = "language="

	BotUnexpectedErrorMsg = "I'm vewy sorwy but somwthing weird happened... >0<"
	BotNoResultsMsg       = "I'm so sowwy... I didn't find anything... (o.O)"
	BotNoActiveStreamsMsg = "Sowwy... There doesn't seem to be a stream like that active at the moment <0.o>"
	BotNoGames            = "...I'm afraid there are no games like that.. sowwy >.<"
)
