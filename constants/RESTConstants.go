package constants


const (
	UrlTwitchChannelId   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	UrlTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	UrlTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
	UrlTwitchStreamInfo  = "https://api.twitch.tv/helix/streams"
	UrlTwitchGames = "https://api.twitch.tv/helix/search/categories?query=" // + [game_name]

	TwitchApiResolution = "52x72.jpg"
	DiscordBotImgResolution = "200x150"

	ParaUserLogin = "user_login="
	ParaGameId = "game_id="
	ParaLanguage = "language="


	BotUnexpectedErrorMessage = "I'm vewy sorwy but somwthing wierd happened... >0<"
	BotNoResults = "I'm so sowwy... I didn't find anything... (o.O)"


)
