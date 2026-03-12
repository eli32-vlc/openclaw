# Agent Pipeline

How OpenClaw spawns, supervises, and manages AI agent runs — from inbound message to final reply.

```mermaid
flowchart TD
    Trigger["Inbound Message\n(channel monitor)"]
    AutoReply["Auto-Reply Engine\nsrc/auto-reply/"]
    Scope["Agent Scope\nsrc/agents/agent-scope.ts"]
    AuthProfile["Auth Profile Resolution\nsrc/agents/auth-profiles/"]
    CLIRunner["CLI Runner\nsrc/agents/cli-runner.ts"]
    CLIHelpers["CLI Helpers\nsrc/agents/cli-runner/helpers.ts"]
    Bootstrap["Bootstrap Files\nsrc/agents/bootstrap-files.ts"]
    BootstrapBudget["Bootstrap Budget\nsrc/agents/bootstrap-budget.ts"]
    SystemPrompt["System Prompt Builder"]
    ProcessSupervisor["Process Supervisor\nsrc/process/supervisor/"]
    CLIBackend["CLI Backend Config\nsrc/agents/cli-backends.ts"]
    FailoverError["Failover Error\nsrc/agents/failover-error.ts"]
    Result["EmbeddedPiRunResult"]
    ReplyDispatcher["Reply Dispatcher\nsrc/auto-reply/reply/reply-dispatcher.ts"]

    Trigger --> AutoReply
    AutoReply --> Scope
    Scope --> AuthProfile
    AuthProfile --> CLIRunner
    CLIRunner --> CLIHelpers
    CLIRunner --> Bootstrap
    Bootstrap --> BootstrapBudget
    CLIHelpers --> SystemPrompt
    CLIHelpers --> CLIBackend
    CLIRunner --> ProcessSupervisor
    ProcessSupervisor --> Result
    Result --> FailoverError
    FailoverError --> ReplyDispatcher
    ReplyDispatcher --> Trigger
```

## Auth Profile Resolution

```mermaid
flowchart LR
    A["resolveAuthProfileOrder()"] --> B{"Profile eligible?"}
    B -- yes --> C["resolveApiKeyForProfile()"]
    B -- no --> D["try next profile"]
    C --> E{"OAuth?\nAPI Key?"}
    E -- OAuth --> F["refresh token if needed"]
    E -- APIKey --> G["use stored key"]
    F --> H["credentials ready"]
    G --> H
    D --> A
```

## CLI Backend Providers

```mermaid
graph LR
    CLIBackend["CLI Backend Config"]

    CLIBackend --> ClaudeCLI["claude-cli\n(Anthropic CLI)"]
    CLIBackend --> CodexCLI["codex-cli\n(Codex CLI)"]
    CLIBackend --> PiEmbedded["pi-embedded\n(embedded pi-ai)"]

    ClaudeCLI --> Anthropic["Anthropic API\nor Bedrock/Vertex"]
    CodexCLI --> OpenAI["OpenAI API"]
    PiEmbedded --> AnyModel["Any supported model\n(via provider config)"]
```

## Session and Workspace Layout

```mermaid
graph TD
    SessionId["sessionId\n(unique per conversation)"]
    SessionKey["sessionKey\n(canonical routing key)"]
    AgentId["agentId\n(optional sub-agent)"]
    WorkspaceDir["workspaceDir\n(resolved via config)"]
    BootstrapFiles["bootstrap files\n(AGENTS.md, CLAUDE.md, etc.)"]
    SessionFile["session .jsonl transcript"]

    SessionId --> WorkspaceDir
    SessionKey --> WorkspaceDir
    AgentId --> WorkspaceDir
    WorkspaceDir --> BootstrapFiles
    WorkspaceDir --> SessionFile
```

## Key Files

| File | Role |
|------|------|
| `src/agents/cli-runner.ts` | Main agent entry — builds args, runs process |
| `src/agents/cli-runner/helpers.ts` | System prompt, CLI args, image handling |
| `src/agents/auth-profiles/` | Credential store, OAuth, API-key resolution |
| `src/agents/bootstrap-files.ts` | Bootstrap context file resolution |
| `src/agents/bootstrap-budget.ts` | Token budget analysis for bootstrap |
| `src/agents/cli-backends.ts` | Backend config (claude-cli / codex / pi-embedded) |
| `src/agents/failover-error.ts` | Failover detection and retry logic |
| `src/process/supervisor/` | Child-process spawn, track, cancel |
