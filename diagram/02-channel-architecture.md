# Channel Architecture

How OpenClaw registers, manages, and dispatches messages across all supported messaging channels.

```mermaid
graph TD
    Registry["Channel Registry\nsrc/channels/registry.ts"]
    PluginDock["Channel Dock\nsrc/channels/dock.ts"]

    subgraph CoreChannels["Core Messaging Channels"]
        Telegram["Telegram\nsrc/telegram/"]
        WhatsApp["WhatsApp\nsrc/whatsapp/"]
        Discord["Discord\nsrc/discord/"]
        Slack["Slack\nsrc/slack/"]
        Signal["Signal\nsrc/signal/"]
        iMessage["iMessage\nsrc/imessage/"]
        IRC["IRC\nextensions/"]
        GoogleChat["Google Chat\nextensions/"]
        LINE["LINE\nsrc/line/"]
    end

    subgraph ExtChannels["Extension Channels (plugins)"]
        MSTeams["MS Teams\nextensions/msteams/"]
        Matrix["Matrix\nextensions/matrix/"]
        Zalo["Zalo\nextensions/zalo/"]
        VoiceCall["Voice Call\nextensions/voice-call/"]
    end

    subgraph ChannelInfra["Channel Infrastructure"]
        Session["Session Tracking\nsrc/channels/session.ts"]
        Allowlist["Allowlist Matching\nsrc/channels/allowlist-match.ts"]
        StateMachine["Run State Machine\nsrc/channels/run-state-machine.ts"]
        Typing["Typing Lifecycle\nsrc/channels/typing-lifecycle.ts"]
        DraftStream["Draft Stream Loop\nsrc/channels/draft-stream-loop.ts"]
        Transport["Stall Watchdog\nsrc/channels/transport/"]
    end

    Registry --> PluginDock
    PluginDock --> CoreChannels
    PluginDock --> ExtChannels
    Registry --> ChannelInfra

    Telegram --> Session
    WhatsApp --> Session
    Discord --> Session
    Slack --> Session
    Signal --> Session
    iMessage --> Session

    Session --> Allowlist
    Session --> StateMachine
    StateMachine --> Typing
    StateMachine --> DraftStream
    DraftStream --> Transport
```

## Channel Registration Flow

```mermaid
sequenceDiagram
    participant Gateway
    participant Registry as Channel Registry
    participant Plugin as Channel Plugin
    participant Dock as Channel Dock
    participant Monitor as Channel Monitor

    Gateway->>Registry: requireActivePluginRegistry()
    Registry->>Plugin: load channel plugin
    Plugin->>Dock: registerChannel(channelId, meta)
    Dock-->>Registry: channel registered
    Gateway->>Monitor: startMonitor(config, deps)
    Monitor-->>Gateway: monitoring started (polling / webhook)
```

## Supported Channels

| Channel | Type | Path |
|---------|------|------|
| Telegram | Core | `src/telegram/` |
| WhatsApp | Core | `src/whatsapp/` |
| Discord | Core | `src/discord/` |
| Slack | Core | `src/slack/` |
| Signal | Core | `src/signal/` |
| iMessage | Core | `src/imessage/` |
| LINE | Core | `src/line/` |
| IRC | Core/Ext | `extensions/` |
| Google Chat | Ext | `extensions/` |
| MS Teams | Ext | `extensions/msteams/` |
| Matrix | Ext | `extensions/matrix/` |
| Zalo | Ext | `extensions/zalo/` |
| Voice Call | Ext | `extensions/voice-call/` |
