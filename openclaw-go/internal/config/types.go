package config

// ReplyMode controls how the agent replies.
type ReplyMode string

const (
	ReplyModeText    ReplyMode = "text"
	ReplyModeCommand ReplyMode = "command"
)

// TypingMode controls typing indicator behavior.
type TypingMode string

const (
	TypingModeNever    TypingMode = "never"
	TypingModeInstant  TypingMode = "instant"
	TypingModeThinking TypingMode = "thinking"
	TypingModeMessage  TypingMode = "message"
)

// SessionScope controls session scoping.
type SessionScope string

const (
	SessionScopePerSender SessionScope = "per-sender"
	SessionScopeGlobal    SessionScope = "global"
)

// DmScope controls DM session scoping.
type DmScope string

const (
	DmScopeMain                  DmScope = "main"
	DmScopePerPeer               DmScope = "per-peer"
	DmScopePerChannelPeer        DmScope = "per-channel-peer"
	DmScopePerAccountChannelPeer DmScope = "per-account-channel-peer"
)

// GroupPolicy controls group access.
type GroupPolicy string

const (
	GroupPolicyOpen      GroupPolicy = "open"
	GroupPolicyDisabled  GroupPolicy = "disabled"
	GroupPolicyAllowlist GroupPolicy = "allowlist"
)

// DmPolicy controls DM access.
type DmPolicy string

const (
	DmPolicyPairing   DmPolicy = "pairing"
	DmPolicyAllowlist DmPolicy = "allowlist"
	DmPolicyOpen      DmPolicy = "open"
	DmPolicyDisabled  DmPolicy = "disabled"
)

// ReplyToMode controls reply-to behavior.
type ReplyToMode string

const (
	ReplyToModeOff   ReplyToMode = "off"
	ReplyToModeFirst ReplyToMode = "first"
	ReplyToModeAll   ReplyToMode = "all"
)

// OutboundRetryConfig configures retry behavior for outbound requests.
type OutboundRetryConfig struct {
	Attempts   *int     `json:"attempts,omitempty" yaml:"attempts,omitempty"`
	MinDelayMs *int     `json:"minDelayMs,omitempty" yaml:"minDelayMs,omitempty"`
	MaxDelayMs *int     `json:"maxDelayMs,omitempty" yaml:"maxDelayMs,omitempty"`
	Jitter     *float64 `json:"jitter,omitempty" yaml:"jitter,omitempty"`
}

// BlockStreamingCoalesceConfig configures block streaming coalescing.
type BlockStreamingCoalesceConfig struct {
	MinChars *int `json:"minChars,omitempty" yaml:"minChars,omitempty"`
	MaxChars *int `json:"maxChars,omitempty" yaml:"maxChars,omitempty"`
	IdleMs   *int `json:"idleMs,omitempty" yaml:"idleMs,omitempty"`
}

// BlockStreamingChunkConfig configures block streaming chunking.
type BlockStreamingChunkConfig struct {
	MinChars        *int   `json:"minChars,omitempty" yaml:"minChars,omitempty"`
	MaxChars        *int   `json:"maxChars,omitempty" yaml:"maxChars,omitempty"`
	BreakPreference string `json:"breakPreference,omitempty" yaml:"breakPreference,omitempty"`
}

// MarkdownTableMode controls how markdown tables are rendered.
type MarkdownTableMode string

const (
	MarkdownTableModeOff     MarkdownTableMode = "off"
	MarkdownTableModeBullets MarkdownTableMode = "bullets"
	MarkdownTableModeCode    MarkdownTableMode = "code"
)

// MarkdownConfig configures markdown rendering.
type MarkdownConfig struct {
	Tables MarkdownTableMode `json:"tables,omitempty" yaml:"tables,omitempty"`
}

// HumanDelayConfig configures human-like delay between replies.
type HumanDelayConfig struct {
	Mode  string `json:"mode,omitempty" yaml:"mode,omitempty"`
	MinMs *int   `json:"minMs,omitempty" yaml:"minMs,omitempty"`
	MaxMs *int   `json:"maxMs,omitempty" yaml:"maxMs,omitempty"`
}

// SessionResetConfig configures session reset behavior.
type SessionResetConfig struct {
	Mode        string `json:"mode,omitempty" yaml:"mode,omitempty"`
	AtHour      *int   `json:"atHour,omitempty" yaml:"atHour,omitempty"`
	IdleMinutes *int   `json:"idleMinutes,omitempty" yaml:"idleMinutes,omitempty"`
}

// SessionResetByTypeConfig configures session reset by chat type.
type SessionResetByTypeConfig struct {
	Direct *SessionResetConfig `json:"direct,omitempty" yaml:"direct,omitempty"`
	Dm     *SessionResetConfig `json:"dm,omitempty" yaml:"dm,omitempty"`
	Group  *SessionResetConfig `json:"group,omitempty" yaml:"group,omitempty"`
	Thread *SessionResetConfig `json:"thread,omitempty" yaml:"thread,omitempty"`
}

// SessionSendPolicyConfig configures session send policy.
type SessionSendPolicyConfig struct {
	Default string                  `json:"default,omitempty" yaml:"default,omitempty"`
	Rules   []SessionSendPolicyRule `json:"rules,omitempty" yaml:"rules,omitempty"`
}

// SessionSendPolicyRule is a single rule in the send policy.
type SessionSendPolicyRule struct {
	Action string                  `json:"action" yaml:"action"`
	Match  *SessionSendPolicyMatch `json:"match,omitempty" yaml:"match,omitempty"`
}

// SessionSendPolicyMatch defines matching criteria for a send policy rule.
type SessionSendPolicyMatch struct {
	Channel      string `json:"channel,omitempty" yaml:"channel,omitempty"`
	ChatType     string `json:"chatType,omitempty" yaml:"chatType,omitempty"`
	KeyPrefix    string `json:"keyPrefix,omitempty" yaml:"keyPrefix,omitempty"`
	RawKeyPrefix string `json:"rawKeyPrefix,omitempty" yaml:"rawKeyPrefix,omitempty"`
}

// SessionThreadBindingsConfig configures thread-bound session routing.
type SessionThreadBindingsConfig struct {
	Enabled     *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	IdleHours   *int  `json:"idleHours,omitempty" yaml:"idleHours,omitempty"`
	MaxAgeHours *int  `json:"maxAgeHours,omitempty" yaml:"maxAgeHours,omitempty"`
}

// SessionMaintenanceConfig configures session store maintenance.
type SessionMaintenanceConfig struct {
	Mode                  string      `json:"mode,omitempty" yaml:"mode,omitempty"`
	PruneAfter            interface{} `json:"pruneAfter,omitempty" yaml:"pruneAfter,omitempty"`
	PruneDays             *int        `json:"pruneDays,omitempty" yaml:"pruneDays,omitempty"`
	MaxEntries            *int        `json:"maxEntries,omitempty" yaml:"maxEntries,omitempty"`
	RotateBytes           interface{} `json:"rotateBytes,omitempty" yaml:"rotateBytes,omitempty"`
	ResetArchiveRetention interface{} `json:"resetArchiveRetention,omitempty" yaml:"resetArchiveRetention,omitempty"`
	MaxDiskBytes          interface{} `json:"maxDiskBytes,omitempty" yaml:"maxDiskBytes,omitempty"`
}

// SessionConfig configures session behavior.
type SessionConfig struct {
	Scope                 SessionScope                   `json:"scope,omitempty" yaml:"scope,omitempty"`
	DmScope               DmScope                        `json:"dmScope,omitempty" yaml:"dmScope,omitempty"`
	IdentityLinks         map[string][]string            `json:"identityLinks,omitempty" yaml:"identityLinks,omitempty"`
	ResetTriggers         []string                       `json:"resetTriggers,omitempty" yaml:"resetTriggers,omitempty"`
	IdleMinutes           *int                           `json:"idleMinutes,omitempty" yaml:"idleMinutes,omitempty"`
	Reset                 *SessionResetConfig            `json:"reset,omitempty" yaml:"reset,omitempty"`
	ResetByType           *SessionResetByTypeConfig      `json:"resetByType,omitempty" yaml:"resetByType,omitempty"`
	ResetByChannel        map[string]*SessionResetConfig `json:"resetByChannel,omitempty" yaml:"resetByChannel,omitempty"`
	Store                 string                         `json:"store,omitempty" yaml:"store,omitempty"`
	TypingIntervalSeconds *int                           `json:"typingIntervalSeconds,omitempty" yaml:"typingIntervalSeconds,omitempty"`
	TypingMode            TypingMode                     `json:"typingMode,omitempty" yaml:"typingMode,omitempty"`
	ParentForkMaxTokens   *int                           `json:"parentForkMaxTokens,omitempty" yaml:"parentForkMaxTokens,omitempty"`
	MainKey               string                         `json:"mainKey,omitempty" yaml:"mainKey,omitempty"`
	SendPolicy            *SessionSendPolicyConfig       `json:"sendPolicy,omitempty" yaml:"sendPolicy,omitempty"`
	ThreadBindings        *SessionThreadBindingsConfig   `json:"threadBindings,omitempty" yaml:"threadBindings,omitempty"`
	Maintenance           *SessionMaintenanceConfig      `json:"maintenance,omitempty" yaml:"maintenance,omitempty"`
}

// LoggingConfig configures logging.
type LoggingConfig struct {
	Level    string `json:"level,omitempty" yaml:"level,omitempty"`
	Console  *bool  `json:"console,omitempty" yaml:"console,omitempty"`
	File     *bool  `json:"file,omitempty" yaml:"file,omitempty"`
	FilePath string `json:"filePath,omitempty" yaml:"filePath,omitempty"`
}

// DiagnosticsConfig configures diagnostics.
type DiagnosticsConfig struct {
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

// TelegramAccountConfig configures a single Telegram bot account.
type TelegramAccountConfig struct {
	Token       string                 `json:"token,omitempty" yaml:"token,omitempty"`
	AllowFrom   []string               `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy    DmPolicy               `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	GroupPolicy GroupPolicy            `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	Webhook     *TelegramWebhookConfig `json:"webhook,omitempty" yaml:"webhook,omitempty"`
}

// TelegramWebhookConfig configures Telegram webhook.
type TelegramWebhookConfig struct {
	Enabled *bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	URL     string `json:"url,omitempty" yaml:"url,omitempty"`
	Secret  string `json:"secret,omitempty" yaml:"secret,omitempty"`
}

// TelegramConfig configures the Telegram channel.
type TelegramConfig struct {
	Enabled        *bool                             `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Token          string                            `json:"token,omitempty" yaml:"token,omitempty"`
	AllowFrom      []string                          `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy       DmPolicy                          `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	GroupPolicy    GroupPolicy                       `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	DefaultAccount string                            `json:"defaultAccount,omitempty" yaml:"defaultAccount,omitempty"`
	Accounts       map[string]*TelegramAccountConfig `json:"accounts,omitempty" yaml:"accounts,omitempty"`
}

// WhatsAppConfig configures the WhatsApp channel.
type WhatsAppConfig struct {
	Enabled        *bool       `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	AllowFrom      []string    `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy       DmPolicy    `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	GroupPolicy    GroupPolicy `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	DefaultAccount string      `json:"defaultAccount,omitempty" yaml:"defaultAccount,omitempty"`
	DefaultTo      string      `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// DiscordConfig configures the Discord channel.
type DiscordConfig struct {
	Enabled        *bool       `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Token          string      `json:"token,omitempty" yaml:"token,omitempty"`
	ClientID       string      `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	AllowFrom      []string    `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy       DmPolicy    `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	GroupPolicy    GroupPolicy `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	DefaultAccount string      `json:"defaultAccount,omitempty" yaml:"defaultAccount,omitempty"`
	DefaultTo      string      `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// SlackConfig configures the Slack channel.
type SlackConfig struct {
	Enabled     *bool       `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	BotToken    string      `json:"botToken,omitempty" yaml:"botToken,omitempty"`
	AppToken    string      `json:"appToken,omitempty" yaml:"appToken,omitempty"`
	AllowFrom   []string    `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy    DmPolicy    `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	GroupPolicy GroupPolicy `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	DefaultTo   string      `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// SignalConfig configures the Signal channel.
type SignalConfig struct {
	Enabled     *bool       `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	PhoneNumber string      `json:"phoneNumber,omitempty" yaml:"phoneNumber,omitempty"`
	ApiURL      string      `json:"apiUrl,omitempty" yaml:"apiUrl,omitempty"`
	AllowFrom   []string    `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy    DmPolicy    `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	GroupPolicy GroupPolicy `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	DefaultTo   string      `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// IMessageConfig configures the iMessage channel.
type IMessageConfig struct {
	Enabled   *bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	AllowFrom []string `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy  DmPolicy `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	DefaultTo string   `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// IRCConfig configures the IRC channel.
type IRCConfig struct {
	Enabled   *bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Server    string   `json:"server,omitempty" yaml:"server,omitempty"`
	Port      *int     `json:"port,omitempty" yaml:"port,omitempty"`
	Nick      string   `json:"nick,omitempty" yaml:"nick,omitempty"`
	Password  string   `json:"password,omitempty" yaml:"password,omitempty"`
	Channels  []string `json:"channels,omitempty" yaml:"channels,omitempty"`
	AllowFrom []string `json:"allowFrom,omitempty" yaml:"allowFrom,omitempty"`
	DmPolicy  DmPolicy `json:"dmPolicy,omitempty" yaml:"dmPolicy,omitempty"`
	TLS       *bool    `json:"tls,omitempty" yaml:"tls,omitempty"`
}

// GoogleChatConfig configures the Google Chat channel.
type GoogleChatConfig struct {
	Enabled           *bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ServiceAccountKey string `json:"serviceAccountKey,omitempty" yaml:"serviceAccountKey,omitempty"`
	ProjectID         string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	DefaultTo         string `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// MSTeamsConfig configures the MS Teams channel.
type MSTeamsConfig struct {
	Enabled      *bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	TenantID     string `json:"tenantId,omitempty" yaml:"tenantId,omitempty"`
	ClientID     string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`
	DefaultTo    string `json:"defaultTo,omitempty" yaml:"defaultTo,omitempty"`
}

// ChannelHeartbeatVisibilityConfig configures channel heartbeat visibility.
type ChannelHeartbeatVisibilityConfig struct {
	ShowOk       *bool `json:"showOk,omitempty" yaml:"showOk,omitempty"`
	ShowAlerts   *bool `json:"showAlerts,omitempty" yaml:"showAlerts,omitempty"`
	UseIndicator *bool `json:"useIndicator,omitempty" yaml:"useIndicator,omitempty"`
}

// ChannelDefaultsConfig configures channel defaults.
type ChannelDefaultsConfig struct {
	GroupPolicy GroupPolicy                       `json:"groupPolicy,omitempty" yaml:"groupPolicy,omitempty"`
	Heartbeat   *ChannelHeartbeatVisibilityConfig `json:"heartbeat,omitempty" yaml:"heartbeat,omitempty"`
}

// ChannelsConfig configures all channels.
type ChannelsConfig struct {
	Defaults       *ChannelDefaultsConfig       `json:"defaults,omitempty" yaml:"defaults,omitempty"`
	ModelByChannel map[string]map[string]string `json:"modelByChannel,omitempty" yaml:"modelByChannel,omitempty"`
	WhatsApp       *WhatsAppConfig              `json:"whatsapp,omitempty" yaml:"whatsapp,omitempty"`
	Telegram       *TelegramConfig              `json:"telegram,omitempty" yaml:"telegram,omitempty"`
	Discord        *DiscordConfig               `json:"discord,omitempty" yaml:"discord,omitempty"`
	IRC            *IRCConfig                   `json:"irc,omitempty" yaml:"irc,omitempty"`
	GoogleChat     *GoogleChatConfig            `json:"googlechat,omitempty" yaml:"googlechat,omitempty"`
	Slack          *SlackConfig                 `json:"slack,omitempty" yaml:"slack,omitempty"`
	Signal         *SignalConfig                `json:"signal,omitempty" yaml:"signal,omitempty"`
	IMessage       *IMessageConfig              `json:"imessage,omitempty" yaml:"imessage,omitempty"`
	MSTeams        *MSTeamsConfig               `json:"msteams,omitempty" yaml:"msteams,omitempty"`
}

// AuthProfileConfig configures a single auth profile.
type AuthProfileConfig struct {
	Provider string `json:"provider" yaml:"provider"`
	Mode     string `json:"mode" yaml:"mode"`
	Email    string `json:"email,omitempty" yaml:"email,omitempty"`
}

// AuthCooldownsConfig configures auth cooldowns.
type AuthCooldownsConfig struct {
	BillingBackoffHours           *int           `json:"billingBackoffHours,omitempty" yaml:"billingBackoffHours,omitempty"`
	BillingBackoffHoursByProvider map[string]int `json:"billingBackoffHoursByProvider,omitempty" yaml:"billingBackoffHoursByProvider,omitempty"`
	BillingMaxHours               *int           `json:"billingMaxHours,omitempty" yaml:"billingMaxHours,omitempty"`
	FailureWindowHours            *int           `json:"failureWindowHours,omitempty" yaml:"failureWindowHours,omitempty"`
}

// AuthConfig configures authentication.
type AuthConfig struct {
	Profiles  map[string]*AuthProfileConfig `json:"profiles,omitempty" yaml:"profiles,omitempty"`
	Order     map[string][]string           `json:"order,omitempty" yaml:"order,omitempty"`
	Cooldowns *AuthCooldownsConfig          `json:"cooldowns,omitempty" yaml:"cooldowns,omitempty"`
}

// GatewayBindMode controls how the gateway binds.
type GatewayBindMode string

const (
	GatewayBindModeAuto     GatewayBindMode = "auto"
	GatewayBindModeLAN      GatewayBindMode = "lan"
	GatewayBindModeLoopback GatewayBindMode = "loopback"
	GatewayBindModeCustom   GatewayBindMode = "custom"
	GatewayBindModeTailnet  GatewayBindMode = "tailnet"
)

// GatewayTlsConfig configures TLS for the gateway.
type GatewayTlsConfig struct {
	Enabled      *bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	AutoGenerate *bool  `json:"autoGenerate,omitempty" yaml:"autoGenerate,omitempty"`
	CertPath     string `json:"certPath,omitempty" yaml:"certPath,omitempty"`
	KeyPath      string `json:"keyPath,omitempty" yaml:"keyPath,omitempty"`
	CaPath       string `json:"caPath,omitempty" yaml:"caPath,omitempty"`
}

// GatewayConfig configures the gateway server.
type GatewayConfig struct {
	Mode      GatewayBindMode   `json:"mode,omitempty" yaml:"mode,omitempty"`
	Host      string            `json:"host,omitempty" yaml:"host,omitempty"`
	Token     string            `json:"token,omitempty" yaml:"token,omitempty"`
	Password  string            `json:"password,omitempty" yaml:"password,omitempty"`
	TLS       *GatewayTlsConfig `json:"tls,omitempty" yaml:"tls,omitempty"`
	ControlUI *bool             `json:"controlUi,omitempty" yaml:"controlUi,omitempty"`
	Port      *int              `json:"port,omitempty" yaml:"port,omitempty"`
}

// AgentModelConfig configures the model for an agent.
type AgentModelConfig struct {
	Provider  string   `json:"provider,omitempty" yaml:"provider,omitempty"`
	Model     string   `json:"model,omitempty" yaml:"model,omitempty"`
	Fallbacks []string `json:"fallbacks,omitempty" yaml:"fallbacks,omitempty"`
}

// AgentBindingMatch defines matching criteria for agent routing.
type AgentBindingMatch struct {
	Channel   string `json:"channel" yaml:"channel"`
	AccountID string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
	GuildID   string `json:"guildId,omitempty" yaml:"guildId,omitempty"`
	TeamID    string `json:"teamId,omitempty" yaml:"teamId,omitempty"`
}

// AgentBinding binds an agent to a channel/match.
type AgentBinding struct {
	Type    string            `json:"type,omitempty" yaml:"type,omitempty"`
	AgentID string            `json:"agentId" yaml:"agentId"`
	Comment string            `json:"comment,omitempty" yaml:"comment,omitempty"`
	Match   AgentBindingMatch `json:"match" yaml:"match"`
}

// AgentConfig configures a single agent.
type AgentConfig struct {
	ID         string            `json:"id" yaml:"id"`
	Default    *bool             `json:"default,omitempty" yaml:"default,omitempty"`
	Name       string            `json:"name,omitempty" yaml:"name,omitempty"`
	Workspace  string            `json:"workspace,omitempty" yaml:"workspace,omitempty"`
	AgentDir   string            `json:"agentDir,omitempty" yaml:"agentDir,omitempty"`
	Model      *AgentModelConfig `json:"model,omitempty" yaml:"model,omitempty"`
	Skills     []string          `json:"skills,omitempty" yaml:"skills,omitempty"`
	HumanDelay *HumanDelayConfig `json:"humanDelay,omitempty" yaml:"humanDelay,omitempty"`
}

// AgentsConfig configures agents.
type AgentsConfig struct {
	Bindings []AgentBinding `json:"bindings,omitempty" yaml:"bindings,omitempty"`
	Agents   []AgentConfig  `json:"agents,omitempty" yaml:"agents,omitempty"`
}

// ModelsConfig configures models.
type ModelsConfig struct {
	Default   string            `json:"default,omitempty" yaml:"default,omitempty"`
	Overrides map[string]string `json:"overrides,omitempty" yaml:"overrides,omitempty"`
}

// SecretsConfig configures secrets storage.
type SecretsConfig struct {
	Store string `json:"store,omitempty" yaml:"store,omitempty"`
}

// PluginConfig configures a single plugin.
type PluginConfig struct {
	Enabled *bool                  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}

// PluginsConfig configures plugins.
type PluginsConfig struct {
	Plugins map[string]*PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// SkillConfig configures a single skill.
type SkillConfig struct {
	Enabled *bool                  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}

// SkillsConfig configures skills.
type SkillsConfig struct {
	Skills map[string]*SkillConfig `json:"skills,omitempty" yaml:"skills,omitempty"`
}

// MemoryConfig configures memory storage.
type MemoryConfig struct {
	Enabled  *bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Provider string `json:"provider,omitempty" yaml:"provider,omitempty"`
	Path     string `json:"path,omitempty" yaml:"path,omitempty"`
}

// CronEntry configures a single cron job.
type CronEntry struct {
	ID       string `json:"id,omitempty" yaml:"id,omitempty"`
	Schedule string `json:"schedule" yaml:"schedule"`
	Message  string `json:"message,omitempty" yaml:"message,omitempty"`
	AgentID  string `json:"agentId,omitempty" yaml:"agentId,omitempty"`
	Enabled  *bool  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

// CronConfig configures cron jobs.
type CronConfig struct {
	Jobs []CronEntry `json:"jobs,omitempty" yaml:"jobs,omitempty"`
}

// HookConfig configures a single hook.
type HookConfig struct {
	Enabled *bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	URL     string   `json:"url,omitempty" yaml:"url,omitempty"`
	Secret  string   `json:"secret,omitempty" yaml:"secret,omitempty"`
	Events  []string `json:"events,omitempty" yaml:"events,omitempty"`
}

// HooksConfig configures hooks.
type HooksConfig struct {
	Hooks map[string]*HookConfig `json:"hooks,omitempty" yaml:"hooks,omitempty"`
}

// ToolConfig configures a single tool.
type ToolConfig struct {
	Enabled *bool                  `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Config  map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}

// ToolsConfig configures tools.
type ToolsConfig struct {
	Tools map[string]*ToolConfig `json:"tools,omitempty" yaml:"tools,omitempty"`
}

// OpenClawConfig is the main configuration struct.
type OpenClawConfig struct {
	Version     string             `json:"version,omitempty" yaml:"version,omitempty"`
	Channels    *ChannelsConfig    `json:"channels,omitempty" yaml:"channels,omitempty"`
	Gateway     *GatewayConfig     `json:"gateway,omitempty" yaml:"gateway,omitempty"`
	Auth        *AuthConfig        `json:"auth,omitempty" yaml:"auth,omitempty"`
	Agents      *AgentsConfig      `json:"agents,omitempty" yaml:"agents,omitempty"`
	Session     *SessionConfig     `json:"session,omitempty" yaml:"session,omitempty"`
	Logging     *LoggingConfig     `json:"logging,omitempty" yaml:"logging,omitempty"`
	Diagnostics *DiagnosticsConfig `json:"diagnostics,omitempty" yaml:"diagnostics,omitempty"`
	Models      *ModelsConfig      `json:"models,omitempty" yaml:"models,omitempty"`
	Secrets     *SecretsConfig     `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	Plugins     *PluginsConfig     `json:"plugins,omitempty" yaml:"plugins,omitempty"`
	Skills      *SkillsConfig      `json:"skills,omitempty" yaml:"skills,omitempty"`
	Memory      *MemoryConfig      `json:"memory,omitempty" yaml:"memory,omitempty"`
	Cron        *CronConfig        `json:"cron,omitempty" yaml:"cron,omitempty"`
	Hooks       *HooksConfig       `json:"hooks,omitempty" yaml:"hooks,omitempty"`
	Tools       *ToolsConfig       `json:"tools,omitempty" yaml:"tools,omitempty"`
}
