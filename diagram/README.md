# OpenClaw Architecture Diagrams

This folder contains [Mermaid](https://mermaid.js.org/) diagrams that document every major subsystem of the OpenClaw codebase. Each diagram file is self-contained Markdown with embedded Mermaid code blocks.

> Diagrams render automatically on GitHub, in VS Code (Mermaid extension), and on any Mermaid-aware viewer.

## Diagram Index

| File | Subsystem | What it covers |
|------|-----------|----------------|
| [01-system-overview.md](./01-system-overview.md) | **System Overview** | Top-level architecture — entry, CLI, gateway, channels, agents, and all supporting subsystems |
| [02-channel-architecture.md](./02-channel-architecture.md) | **Channels** | All messaging channels (Telegram, WhatsApp, Discord, Slack, Signal, iMessage, LINE, IRC, Matrix, Zalo, MS Teams, Voice), channel registry, dock, session, allowlist, state machine |
| [03-agent-pipeline.md](./03-agent-pipeline.md) | **Agent Pipeline** | End-to-end agent run — auth profile resolution, CLI runner, backend selection (claude-cli / codex / pi-embedded), process supervisor, session/workspace layout |
| [04-config-system.md](./04-config-system.md) | **Configuration** | JSON5 config load/parse/validate/write cycle, Zod schema hierarchy, TypeScript type tree, env-variable substitution, legacy migration |
| [05-gateway-server.md](./05-gateway-server.md) | **Gateway Server** | HTTP gateway server, bind modes, all API methods (config, agents, secrets, skills, nodes, doctor, web, update), node pairing flow |
| [06-message-flow.md](./06-message-flow.md) | **Message Flow** | Full inbound-to-reply sequence: allowlist → session record → debounce → auto-reply → agent run → reply dispatcher → channel send |
| [07-plugin-system.md](./07-plugin-system.md) | **Plugin System** | Plugin discovery, loading, schema validation, registry, API surface (hooks / channels / tools / CLI / services / providers / HTTP routes), lifecycle state machine |
| [08-media-pipeline.md](./08-media-pipeline.md) | **Media Pipeline** | Inbound/outbound image, audio (FFmpeg), video, and PDF processing; MIME sniffing, base64, temp files, media HTTP server |
| [09-memory-system.md](./09-memory-system.md) | **Memory / Embeddings** | SQLite-backed hybrid vector + FTS5 search; embedding providers (OpenAI, Mistral, Ollama, Voyage); MMR re-ranking; QMD query language |
| [10-process-supervision.md](./10-process-supervision.md) | **Process Supervision** | Child-process supervisor state machine, run registry, PTY/child adapters, process lanes, command queue, kill-tree, termination reasons |
| [11-auth-profiles.md](./11-auth-profiles.md) | **Auth Profiles** | API key, OAuth, and token credential types; profile ordering and eligibility; cooldown/failure tracking; external CLI sync; repair and doctor |
| [12-routing-sessions.md](./12-routing-sessions.md) | **Routing & Sessions** | Session key derivation, account/route resolution, explicit bindings, session metadata persistence, ACP persistent bindings, session continuity sequence |

## How to View

- **GitHub**: diagrams render inline in any `.md` file.
- **VS Code**: install the [Markdown Preview Mermaid Support](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-mermaid) extension.
- **CLI**: `npx @mermaid-js/mermaid-cli -i diagram/01-system-overview.md`
- **Online**: paste diagram code into [mermaid.live](https://mermaid.live).
