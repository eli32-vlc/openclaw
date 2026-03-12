# Routing and Sessions

How OpenClaw determines which account and AI session to route an inbound message to, and how session state is persisted.

```mermaid
flowchart TD
    InboundMsg["Inbound Message\n(channel event)"]
    SessionKey["Session Key\nsrc/routing/session-key.ts"]
    AccountLookup["Account Lookup\nsrc/routing/account-lookup.ts"]
    AccountId["Account ID Resolution\nsrc/routing/account-id.ts"]
    ResolveRoute["Resolve Route\nsrc/routing/resolve-route.ts"]
    Bindings["Route Bindings\nsrc/routing/bindings.ts"]
    DefaultAccountWarnings["Default Account Warnings\nsrc/routing/default-account-warnings.ts"]
    ConfigSessions["Session Store\nsrc/config/sessions.ts"]
    SessionMeta["Session Metadata"]
    AgentRun["Agent Run\n(resolved account + session)"]

    InboundMsg --> SessionKey
    SessionKey --> AccountLookup
    AccountLookup --> AccountId
    AccountId --> ResolveRoute
    ResolveRoute --> Bindings
    Bindings --> DefaultAccountWarnings
    ResolveRoute --> ConfigSessions
    ConfigSessions --> SessionMeta
    SessionMeta --> AgentRun
```

## Session Key Derivation

```mermaid
flowchart LR
    Channel["channel ID\n(e.g. telegram)"]
    Sender["sender identifier\n(user ID / phone / username)"]
    ThreadId["thread/topic ID\n(optional, for groups)"]
    AccountId["account ID\n(optional, for multi-account)"]

    Channel --> SessionKey["session key\n(canonical string)"]
    Sender --> SessionKey
    ThreadId --> SessionKey
    AccountId --> SessionKey

    SessionKey --> Normalize["normalize:\ntrim + lowercase"]
```

## Route Resolution

```mermaid
flowchart TD
    RouteReq["Route Request\n(sessionKey, channel, accountId)"]
    ExplicitBinding["Check explicit binding\n(routing.bindings[])"]
    DefaultAccount["Resolve default account\nfor channel"]
    AccountStore["Account Store\n(configured accounts)"]
    LastRoute["Last-used route\n(session store)"]
    FinalRoute["Resolved Route\n(account, session, channel)"]

    RouteReq --> ExplicitBinding
    ExplicitBinding -- found --> FinalRoute
    ExplicitBinding -- not found --> DefaultAccount
    DefaultAccount --> AccountStore
    AccountStore --> LastRoute
    LastRoute --> FinalRoute
```

## Session Metadata Model

```mermaid
classDiagram
    class SessionEntry {
        +sessionKey: string
        +lastChannel: string
        +lastAccountId?: string
        +lastThreadId?: string
        +lastActiveAt: number
        +createdAt: number
        +meta: SessionMeta
    }
    class SessionMeta {
        +displayName?: string
        +avatarUrl?: string
        +groupName?: string
        +isGroup: boolean
    }
    class GroupKeyResolution {
        +groupKey: string
        +mainKey: string
        +threadKey?: string
    }

    SessionEntry --> SessionMeta
    SessionEntry --> GroupKeyResolution
```

## ACP Persistent Bindings

```mermaid
flowchart LR
    ACPBinding["ACP Persistent Binding\nsrc/acp/persistent-bindings.ts"]
    BindingLifecycle["Binding Lifecycle\nsrc/acp/persistent-bindings.lifecycle.ts"]
    BindingResolve["Binding Resolve\nsrc/acp/persistent-bindings.resolve.ts"]
    BindingRoute["Binding Route\nsrc/acp/persistent-bindings.route.ts"]
    BindingTypes["Binding Types\nsrc/acp/persistent-bindings.types.ts"]

    ACPBinding --> BindingLifecycle
    ACPBinding --> BindingResolve
    ACPBinding --> BindingRoute
    ACPBinding --> BindingTypes
    BindingRoute --> ResolveRoute["resolve-route.ts"]
```

## Session Continuity

```mermaid
sequenceDiagram
    participant Monitor as Channel Monitor
    participant SessionTS as channels/session.ts
    participant ConfigSessions as config/sessions.ts
    participant Router as routing/resolve-route.ts

    Monitor->>SessionTS: recordInboundSession(params)
    SessionTS->>ConfigSessions: recordSessionMetaFromInbound()
    SessionTS->>ConfigSessions: updateLastRoute()
    ConfigSessions-->>SessionTS: updated
    Monitor->>Router: resolveRoute(sessionKey, channel)
    Router->>ConfigSessions: getSessionEntry(sessionKey)
    ConfigSessions-->>Router: SessionEntry (lastChannel, lastAccountId)
    Router-->>Monitor: ResolvedRoute
```

## Key Files

| File | Role |
|------|------|
| `src/routing/session-key.ts` | Derive and normalize session keys |
| `src/routing/account-lookup.ts` | Find account config for a session |
| `src/routing/account-id.ts` | Parse and resolve account identifiers |
| `src/routing/resolve-route.ts` | Main route resolution logic |
| `src/routing/bindings.ts` | Static routing bindings from config |
| `src/routing/default-account-warnings.ts` | Warn when default account is ambiguous |
| `src/channels/session.ts` | Record inbound session and last-route updates |
| `src/config/sessions.ts` | Persistent session metadata store |
| `src/acp/persistent-bindings.ts` | ACP-level persistent session bindings |
