package discord

// discord configs
type discord struct {
	Bots []discordBot `json:"bots,omitempty"`
}

type discordBot struct {
	BotName string           `json:"bot_name,omitempty"`
	BotID   string           `json:"bot_id,omitempty"`
	Config  discordBotConfig `json:"config,omitempty"`
	Servers []discordServer  `json:"servers,omitempty"`
}

type discordBotConfig struct {
	Token  string        `json:"token,omitempty"`
	Game   string        `json:"game,omitempty"`
	DMResp ResponseArray `json:"dm_response,omitempty"`
}

type discordServer struct {
	ServerID    string              `json:"server_id,omitempty"`
	Config      discordServerConfig `json:"config,omitempty"`
	ChanGroups  []ChannelGroup      `json:"channel_groups,omitempty"`
	Permissions []Permission        `json:"permissions,omitempty"`
	Filters     []Filter            `json:"filters,omitempty"`
}

type discordServerConfig struct {
	Prefix   string          `json:"prefix,omitempty"`
	Clear    bool            `json:"clear_commands,omitempty"`
	WebHooks discordWebHooks `json:"web_hooks,omitempty"`
}

type discordWebHooks struct {
	Logs string `json:"logs,omitempty"`
}
