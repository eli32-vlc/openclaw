# Configuration System

How OpenClaw loads, validates, merges, and persists its JSON5 configuration.

```mermaid
flowchart TD
    ConfigFile["~/.openclaw/config.json5\n(on-disk file)"]
    ReadSnapshot["readConfigFileSnapshot()\nsrc/config/io.ts"]
    ParseJson5["parseConfigJson5()"]
    Validate["validateConfigObject()\nsrc/config/validation.ts"]
    ZodSchema["Zod Schema\nsrc/config/zod-schema.*.ts"]
    MigrateLegacy["migrateLegacyConfig()\nsrc/config/legacy-migrate.ts"]
    EnvSubstitution["env variable substitution\nsrc/config/env-substitution.ts"]
    RuntimeSnapshot["getRuntimeConfigSnapshot()\n(in-memory cache)"]
    ConfigTypes["TypeScript types\nsrc/config/types.ts"]
    WriteConfig["writeConfigFile()\n(atomic write)"]
    MergePatch["merge-patch\nsrc/config/merge-patch.ts"]

    ConfigFile --> ReadSnapshot
    ReadSnapshot --> ParseJson5
    ParseJson5 --> MigrateLegacy
    MigrateLegacy --> EnvSubstitution
    EnvSubstitution --> Validate
    Validate --> ZodSchema
    ZodSchema --> ConfigTypes
    Validate --> RuntimeSnapshot
    RuntimeSnapshot -->|refresh handler| ConfigFile

    WriteConfig --> MergePatch
    MergePatch --> ConfigFile
```

## Configuration Type Hierarchy

```mermaid
classDiagram
    class OpenClawConfig {
        +AgentsConfig agents
        +TelegramConfig telegram
        +DiscordConfig discord
        +SlackConfig slack
        +SignalConfig signal
        +iMessageConfig imessage
        +SandboxConfig sandbox
        +AuthConfig auth
        +BrowserConfig browser
    }
    class AgentsConfig {
        +string provider
        +string model
        +string[] bootstrapFiles
        +WorkspaceConfig workspace
        +MemorySearchConfig memory
    }
    class TelegramConfig {
        +string token
        +AllowlistConfig allowlist
        +string responsePrefix
    }
    class DiscordConfig {
        +string token
        +string[] guildIds
        +AllowlistConfig allowlist
    }
    class SlackConfig {
        +string botToken
        +string signingSecret
        +AllowlistConfig allowlist
    }

    OpenClawConfig --> AgentsConfig
    OpenClawConfig --> TelegramConfig
    OpenClawConfig --> DiscordConfig
    OpenClawConfig --> SlackConfig
```

## Config IO Lifecycle

```mermaid
sequenceDiagram
    participant CLI
    participant IO as config/io.ts
    participant Disk as ~/.openclaw/config.json5
    participant Cache as RuntimeSnapshot

    CLI->>IO: loadConfig(paths)
    IO->>Disk: read file
    Disk-->>IO: raw JSON5 string
    IO->>IO: parseConfigJson5()
    IO->>IO: migrateLegacyConfig()
    IO->>IO: validateConfigObject()
    IO-->>Cache: setRuntimeConfigSnapshot()
    Cache-->>CLI: OpenClawConfig

    CLI->>IO: writeConfigFile(patch)
    IO->>IO: merge-patch current config
    IO->>Disk: atomic write (tmp → rename)
    IO-->>Cache: clearConfigCache()
```

## Key Files

| File | Role |
|------|------|
| `src/config/io.ts` | File read/write, snapshot cache |
| `src/config/types.ts` | Master TypeScript config types |
| `src/config/validation.ts` | Zod-based schema validation |
| `src/config/zod-schema.*.ts` | Sub-schemas (approvals, sandbox, secrets) |
| `src/config/legacy-migrate.ts` | Migrate old config shapes |
| `src/config/env-substitution.ts` | `$ENV_VAR` interpolation in config values |
| `src/config/merge-patch.ts` | JSON Merge Patch (RFC 7396) writer |
| `src/config/defaults.ts` | Default config values |
| `src/config/paths.ts` | Resolve config file paths |
