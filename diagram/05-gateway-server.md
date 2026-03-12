# Gateway Server

The OpenClaw gateway is the central HTTP/WS server that bridges channels, agents, config, and remote nodes.

```mermaid
graph TD
    GatewayRun["gateway run\n(CLI command)"]
    GatewayServer["Gateway HTTP Server\nsrc/gateway/server*.ts"]
    Auth["Startup Auth\nsrc/gateway/startup-auth.ts"]
    NetConfig["Net Config\nsrc/gateway/net.ts"]
    Credentials["Credentials\nsrc/gateway/credentials.ts"]
    Hooks["Gateway Hooks\nsrc/gateway/hooks.ts"]
    SessionKey["Server Session Key\nsrc/gateway/server-session-key.ts"]
    Maintenance["Maintenance Mode\nsrc/gateway/server-maintenance.ts"]

    subgraph ServerMethods["Gateway API Methods"]
        MethodConfig["config\nsrc/gateway/server-methods/config.ts"]
        MethodAgents["agents\nsrc/gateway/server-methods/agents.ts"]
        MethodSecrets["secrets\nsrc/gateway/server-methods/secrets.ts"]
        MethodSkills["skills\nsrc/gateway/server-methods/skills.ts"]
        MethodNodes["nodes\nsrc/gateway/server-methods/nodes*.ts"]
        MethodDoctor["doctor\nsrc/gateway/server-methods/doctor.ts"]
        MethodWeb["web\nsrc/gateway/server-methods/web.ts"]
        MethodUpdate["update\nsrc/gateway/server-methods/update.ts"]
    end

    GatewayRun --> GatewayServer
    GatewayServer --> Auth
    GatewayServer --> NetConfig
    GatewayServer --> Credentials
    GatewayServer --> Hooks
    GatewayServer --> SessionKey
    GatewayServer --> Maintenance
    GatewayServer --> ServerMethods
```

## Gateway Request Lifecycle

```mermaid
sequenceDiagram
    participant Client as Remote Node / App
    participant GW as Gateway Server
    participant Auth as Auth Middleware
    participant Handler as Method Handler
    participant Config as Config System

    Client->>GW: HTTP request (e.g. POST /api/config)
    GW->>Auth: validate session key / bearer token
    Auth-->>GW: authorized
    GW->>Handler: route to handler
    Handler->>Config: read / write config
    Config-->>Handler: result
    Handler-->>GW: response payload
    GW-->>Client: JSON response
```

## Node Pairing and Registration

```mermaid
flowchart LR
    NodeCLI["openclaw node register\n(remote device)"]
    PairingCode["pairing code\nsrc/pairing/"]
    GatewayPair["gateway pair endpoint\nsrc/gateway/server-methods/nodes*.ts"]
    NodeStore["node store\n(registered nodes)"]
    Channels["channel routing\n(per-node)"]

    NodeCLI --> PairingCode
    PairingCode --> GatewayPair
    GatewayPair --> NodeStore
    NodeStore --> Channels
```

## Gateway Bind Modes

```mermaid
graph LR
    GW["Gateway Server"]
    GW --> Loopback["loopback\n127.0.0.1 (default)"]
    GW --> LAN["lan\n0.0.0.0 (local network)"]
    GW --> Tailscale["tailscale\n(Tailscale IP only)"]
    GW --> Custom["custom\n(explicit --bind address)"]
```

## Key Files

| File | Role |
|------|------|
| `src/gateway/net.ts` | Bind address / port resolution |
| `src/gateway/startup-auth.ts` | Authentication at startup (API key, OAuth) |
| `src/gateway/credentials.ts` | Credential store for gateway |
| `src/gateway/hooks.ts` | Lifecycle hooks fired around gateway events |
| `src/gateway/server-session-key.ts` | Per-session key generation and validation |
| `src/gateway/server-maintenance.ts` | Maintenance mode flag and middleware |
| `src/gateway/server-methods/config.ts` | Read/write gateway config via API |
| `src/gateway/server-methods/agents.ts` | Query/control running agent sessions |
| `src/gateway/server-methods/secrets.ts` | Encrypted secrets management API |
| `src/gateway/server-methods/nodes*.ts` | Node registration, status, pending |
