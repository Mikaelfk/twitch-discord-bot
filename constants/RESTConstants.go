// Package constants provides often used values
package constants

// Often used constants
const (
<<<<<<< HEAD
	URLTwitchChannelID   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	URLTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	URLTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
	URLTwitchStreamInfo  = "https://api.twitch.tv/helix/streams"
	URLTwitchGames       = "https://api.twitch.tv/helix/search/categories?query=" // + [game_name]
	URLTwitchTopGames    = "https://api.twitch.tv/helix/games/top"
	URLTwitchUserName    = "https://api.twitch.tv/helix/users?login="
	URLTwitchFollowlist  = "https://api.twitch.tv/helix/users/follows?from_id="
	URLTwitchGetTeams    = "https://api.twitch.tv/helix/teams?name=" // + [team name]
=======
	UrlTwitchChannelId   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	UrlTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	UrlTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
	UrlTwitchStreamInfo  = "https://api.twitch.tv/helix/streams"
	UrlTwitchGames       = "https://api.twitch.tv/helix/search/categories?query=" // + [game_name]
	UrlTwitchTopGames    = "https://api.twitch.tv/helix/games/top"
	UrlTwitchUserName    = "https://api.twitch.tv/helix/users?login="
	UrlTwitchFollowlist  = "https://api.twitch.tv/helix/users/follows?from_id="
	UrlTwitchGetTeams    = "https://api.twitch.tv/helix/teams?name="			  // + [team name]
>>>>>>> Added a rest constant for the get twitch team url

	TwitchAPIResolution     = "52x72.jpg"
	DiscordBotImgResolution = "200x150"

	ParaUserLogin = "user_login="
	ParaGameID    = "game_id="
	ParaLanguage  = "language="

	BotUnexpectedErrorMsg = "I'm vewy sorwy but somwthing weird happened... >0<"
	BotNoResultsMsg       = "I'm so sowwy... I didn't find anything... (o.O)"
	BotNoActiveStreamsMsg = "Sowwy... There doesn't seem to be a stream like that active at the moment <0.o>"
	BotNoGames            = "...I'm afraid there are no games like that.. sowwy >.<"
)
