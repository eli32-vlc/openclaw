# Message Flow

End-to-end path of a message from arrival on a channel to the final AI reply delivered back to the user.

```mermaid
sequenceDiagram
    participant User
    participant Channel as Channel Monitor<br/>(e.g. telegram/monitor.ts)
    participant Allowlist as Allowlist Check<br/>channels/allowlist-match.ts
    participant Session as Session Record<br/>channels/session.ts
    participant Debounce as Inbound Debounce<br/>channels/inbound-debounce-policy.ts
    participant AutoReply as Auto-Reply Engine<br/>auto-reply/
    participant Agent as Agent CLI Runner<br/>agents/cli-runner.ts
    participant Dispatcher as Reply Dispatcher<br/>auto-reply/reply/reply-dispatcher.ts
    participant Typing as Typing Lifecycle<br/>channels/typing-lifecycle.ts

    User->>Channel: sends message
    Channel->>Allowlist: check sender allowlist
    Allowlist-->>Channel: allowed / denied
    Channel->>Session: recordInboundSession()
    Channel->>Debounce: apply debounce policy
    Debounce-->>Channel: proceed / hold
    Channel->>AutoReply: trigger reply pipeline
    AutoReply->>Typing: startTyping()
    AutoReply->>Agent: runCliAgent(params)
    Agent-->>AutoReply: EmbeddedPiRunResult (streaming)
    AutoReply->>Dispatcher: dispatch reply blocks
    Dispatcher->>Typing: stopTyping()
    Dispatcher-->>Channel: deliver reply
    Channel-->>User: reply message(s)
```

## Inbound Message Processing

```mermaid
flowchart TD
    Raw["Raw update from platform API"]
    EventHandler["Event Handler\n(channel-specific)"]
    GroupCheck["Group / DM check"]
    AllowlistCheck["Allowlist check\nchannels/allowlist-match.ts"]
    MentionGating["Mention gating\n(group channels)"]
    SessionRecord["Session metadata record\nchannels/session.ts"]
    InboundDebounce["Inbound debounce\nchannels/inbound-debounce-policy.ts"]
    BuildContext["Build MsgContext\nauto-reply/templating.ts"]
    CommandCheck["Native command check\nauto-reply/commands-registry.ts"]
    AgentRun["Agent Run\nagents/cli-runner.ts"]

    Raw --> EventHandler
    EventHandler --> GroupCheck
    GroupCheck --> AllowlistCheck
    AllowlistCheck --> MentionGating
    MentionGating --> SessionRecord
    SessionRecord --> InboundDebounce
    InboundDebounce --> BuildContext
    BuildContext --> CommandCheck
    CommandCheck -- built-in cmd --> DirectReply["Direct reply\n(no agent)"]
    CommandCheck -- agent prompt --> AgentRun
```

## Reply Dispatch and Streaming

```mermaid
flowchart LR
    AgentStream["Agent stdout\n(JSONL stream)"]
    ParseBlocks["parseCliJsonl()\nblock-by-block"]

    subgraph Dispatch["Reply Dispatcher"]
        Normalize["normalizeReplyPayload()"]
        HumanDelay["human delay\n(optional)"]
        ResponsePrefix["response prefix\n(optional)"]
        Deliver["deliver(payload, kind)"]
    end

    AckReactions["ack reactions\nchannels/ack-reactions.ts"]
    StatusReactions["status reactions\nchannels/status-reactions.ts"]
    ChannelSend["channel send API\n(e.g. telegram/send.ts)"]

    AgentStream --> ParseBlocks
    ParseBlocks --> Normalize
    Normalize --> HumanDelay
    HumanDelay --> ResponsePrefix
    ResponsePrefix --> Deliver
    Deliver --> AckReactions
    Deliver --> StatusReactions
    Deliver --> ChannelSend
```

## Key Files

| File | Role |
|------|------|
| `src/channels/allowlist-match.ts` | Evaluate sender against allowlist rules |
| `src/channels/inbound-debounce-policy.ts` | Debounce rapid inbound messages |
| `src/channels/session.ts` | Session metadata recording |
| `src/channels/typing-lifecycle.ts` | Manage typing indicator start/stop |
| `src/auto-reply/` | Trigger evaluation, command dispatch, reply pipeline |
| `src/auto-reply/reply/reply-dispatcher.ts` | Block-by-block delivery with optional delay |
| `src/auto-reply/reply/normalize-reply.ts` | Strip heartbeat, prefix, empty content |
| `src/channels/ack-reactions.ts` | Emoji reactions on message acknowledgement |
| `src/channels/status-reactions.ts` | Status emoji reactions (busy, error, etc.) |
