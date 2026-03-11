package channels

// ChatChannelID is the ID of a chat channel.
type ChatChannelID string

const (
	ChannelTelegram   ChatChannelID = "telegram"
	ChannelWhatsApp   ChatChannelID = "whatsapp"
	ChannelDiscord    ChatChannelID = "discord"
	ChannelIRC        ChatChannelID = "irc"
	ChannelGoogleChat ChatChannelID = "googlechat"
	ChannelSlack      ChatChannelID = "slack"
	ChannelSignal     ChatChannelID = "signal"
	ChannelIMessage   ChatChannelID = "imessage"
	ChannelLine       ChatChannelID = "line"
)

// ChatChannelOrder defines the canonical ordering of channels.
var ChatChannelOrder = []ChatChannelID{
	ChannelTelegram,
	ChannelWhatsApp,
	ChannelDiscord,
	ChannelIRC,
	ChannelGoogleChat,
	ChannelSlack,
	ChannelSignal,
	ChannelIMessage,
	ChannelLine,
}

// CHANNEL_IDS is the list of all channel IDs.
var CHANNEL_IDS = ChatChannelOrder

// ChannelMeta holds metadata about a channel.
type ChannelMeta struct {
	ID                     ChatChannelID `json:"id"`
	Label                  string        `json:"label"`
	SelectionLabel         string        `json:"selectionLabel"`
	DetailLabel            string        `json:"detailLabel"`
	DocsPath               string        `json:"docsPath"`
	DocsLabel              string        `json:"docsLabel"`
	Blurb                  string        `json:"blurb"`
	SystemImage            string        `json:"systemImage"`
	SelectionDocsPrefix    string        `json:"selectionDocsPrefix,omitempty"`
	SelectionDocsOmitLabel bool          `json:"selectionDocsOmitLabel,omitempty"`
	SelectionExtras        []string      `json:"selectionExtras,omitempty"`
}

const websiteURL = "https://openclaw.ai"

// ChatChannelMeta holds metadata for all channels.
var ChatChannelMeta = map[ChatChannelID]ChannelMeta{
	ChannelTelegram: {
		ID:                     ChannelTelegram,
		Label:                  "Telegram",
		SelectionLabel:         "Telegram (Bot API)",
		DetailLabel:            "Telegram Bot",
		DocsPath:               "/channels/telegram",
		DocsLabel:              "telegram",
		Blurb:                  "simplest way to get started — register a bot with @BotFather and get going.",
		SystemImage:            "paperplane",
		SelectionDocsPrefix:    "",
		SelectionDocsOmitLabel: true,
		SelectionExtras:        []string{websiteURL},
	},
	ChannelWhatsApp: {
		ID:             ChannelWhatsApp,
		Label:          "WhatsApp",
		SelectionLabel: "WhatsApp (QR link)",
		DetailLabel:    "WhatsApp Web",
		DocsPath:       "/channels/whatsapp",
		DocsLabel:      "whatsapp",
		Blurb:          "works with your own number; recommend a separate phone + eSIM.",
		SystemImage:    "message",
	},
	ChannelDiscord: {
		ID:             ChannelDiscord,
		Label:          "Discord",
		SelectionLabel: "Discord (Bot API)",
		DetailLabel:    "Discord Bot",
		DocsPath:       "/channels/discord",
		DocsLabel:      "discord",
		Blurb:          "very well supported right now.",
		SystemImage:    "bubble.left.and.bubble.right",
	},
	ChannelIRC: {
		ID:             ChannelIRC,
		Label:          "IRC",
		SelectionLabel: "IRC (Server + Nick)",
		DetailLabel:    "IRC",
		DocsPath:       "/channels/irc",
		DocsLabel:      "irc",
		Blurb:          "classic IRC networks with DM/channel routing and pairing controls.",
		SystemImage:    "network",
	},
	ChannelGoogleChat: {
		ID:             ChannelGoogleChat,
		Label:          "Google Chat",
		SelectionLabel: "Google Chat (Chat API)",
		DetailLabel:    "Google Chat",
		DocsPath:       "/channels/googlechat",
		DocsLabel:      "googlechat",
		Blurb:          "Google Workspace Chat app with HTTP webhook.",
		SystemImage:    "message.badge",
	},
	ChannelSlack: {
		ID:             ChannelSlack,
		Label:          "Slack",
		SelectionLabel: "Slack (Socket Mode)",
		DetailLabel:    "Slack Bot",
		DocsPath:       "/channels/slack",
		DocsLabel:      "slack",
		Blurb:          "supported (Socket Mode).",
		SystemImage:    "number",
	},
	ChannelSignal: {
		ID:             ChannelSignal,
		Label:          "Signal",
		SelectionLabel: "Signal (signal-cli)",
		DetailLabel:    "Signal REST",
		DocsPath:       "/channels/signal",
		DocsLabel:      "signal",
		Blurb:          `signal-cli linked device; more setup (David Reagans: "Hop on Discord.").`,
		SystemImage:    "antenna.radiowaves.left.and.right",
	},
	ChannelIMessage: {
		ID:             ChannelIMessage,
		Label:          "iMessage",
		SelectionLabel: "iMessage (imsg)",
		DetailLabel:    "iMessage",
		DocsPath:       "/channels/imessage",
		DocsLabel:      "imessage",
		Blurb:          "this is still a work in progress.",
		SystemImage:    "message.fill",
	},
	ChannelLine: {
		ID:             ChannelLine,
		Label:          "LINE",
		SelectionLabel: "LINE (Messaging API)",
		DetailLabel:    "LINE Bot",
		DocsPath:       "/channels/line",
		DocsLabel:      "line",
		Blurb:          "LINE Messaging API webhook bot.",
		SystemImage:    "message",
	},
}

// GetChannelMeta returns metadata for a channel by ID.
func GetChannelMeta(id ChatChannelID) (ChannelMeta, bool) {
	meta, ok := ChatChannelMeta[id]
	return meta, ok
}
