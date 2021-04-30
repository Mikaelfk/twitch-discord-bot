package constants


const (
	UrlTwitchChannelId   = "https://api.twitch.tv/helix/channels?broadcaster_id=" // + [id]
	UrlTwitchChannelName = "https://api.twitch.tv/helix/search/channels?query="   // + [broadcaster_name]
	UrlTwitchStream      = "https://www.twitch.tv/"                               // + [broadcaster_name]
	UrlTwitchStreamInfo  = "https://api.twitch.tv/helix/streams"
	UrlTwitchGames = "https://api.twitch.tv/helix/search/categories?query=" // + [game_name]


	ParaUserLogin = "user_login="
	ParaGameId = "game_id="
	ParaLanguage = "language="
)
