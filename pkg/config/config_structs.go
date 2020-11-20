package config

type Parkertron struct {
	Services []string       `json:"services,omitempty"`
	Log      LogConf        `json:"log,omitempty"`
	Database DatabaseConfig `json:"database,omitempty"`
	Parsing  BotParseConfig `json:"parsing,omitempty"`
}

type LogConf struct {
	Level    string `json:"level,omitempty"`
	Location string `json:"location,omitempty"`
}

type DatabaseConfig struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Pass     string `json:"pass,omitempty"`
	Database string `json:"database,omitempty"`
}

type BotParseConfig struct {
	Reaction []string `json:"reaction,omitempty"`
	Response []string `json:"response,omitempty"`
	Max      int      `json:"max,omitempty"`
	AllowIP  bool     `json:"allow_ip,omitempty"`
}

// generic structs
type Permission struct {
	Group       string   `json:"group,omitempty"`
	Users       []string `json:"users,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Commands    []string `json:"commands,omitempty"`
	Blacklisted bool     `json:"blacklisted,omitempty"`
}

type Command struct {
	Command  string   `json:"command,omitempty"`
	Response []string `json:"response,omitempty"`
	Reaction []string `json:"reaction,omitempty"`
}

type Keyword struct {
	Keyword  string   `json:"keyword,omitempty"`
	Reaction []string `json:"reaction,omitempty"`
	Response []string `json:"response,omitempty"`
	Exact    bool     `json:"exact,omitempty"`
}

type Mentions struct {
	Ping    ResponseArray `json:"ping,omitempty"`
	Mention ResponseArray `json:"mention,omitempty"`
}

type Filter struct {
	Term   string   `json:"term,omitempty"`
	Reason []string `json:"reason,omitempty"`
}

type ResponseArray struct {
	Reaction []string `json:"reaction,omitempty"`
	Response []string `json:"response,omitempty"`
}

type Parsing struct {
	Image ParsingImageConfig `json:"image,omitempty"`
	Paste ParsingPasteConfig `json:"paste,omitempty"`
}

type ParsingConfig struct {
	Name   string `json:"name,omitempty"`
	URL    string `json:"url,omitempty"`
	Format string `json:"format,omitempty"`
}

type ParsingImageConfig struct {
	FileTypes []string        `json:"filetypes,omitempty"`
	Sites     []ParsingConfig `json:"sites,omitempty"`
}

type ParsingPasteConfig struct {
	Sites  []ParsingConfig `json:"sites,omitempty"`
	Ignore []ParsingConfig `json:"ignore,omitmepty"`
}

type ChannelGroup struct {
	ChannelIDs  []string      `json:"channels,omitempty"`
	Mentions    Mentions      `json:"mentions,omitempty"`
	Commands    []Command     `json:"commands,omitempty"`
	Keywords    []Keyword     `json:"keywords,omitempty"`
	Parsing     Parsing       `json:"parsing,omitempty"`
	Permissions []Permission  `json:"permissions,omitempty"`
	KOM         KickOnMention `json:"kick_on_mention,omitempty"`
}

type KickOnMention struct {
	Roles   []string      `json:"roles,omitempty"`
	Users   []string      `json:"users,omitempty"`
	Direct  ResponseArray `json:"dm,omitempty"`
	Channel ResponseArray `json:"channel,omitempty"`
	Kick    bool          `json:"kick,omitempty"`
}
