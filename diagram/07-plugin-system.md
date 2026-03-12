# Plugin System

How OpenClaw discovers, loads, validates, and runs first-party and third-party plugins (extensions).

```mermaid
flowchart TD
    PluginConfig["Plugin config\n(config.plugins[])"]
    Discovery["Plugin Discovery\nsrc/plugins/discovery.ts"]
    BundledSources["Bundled Sources\nsrc/plugins/bundled-sources.ts"]
    Loader["Plugin Loader\nsrc/plugins/loader.ts"]
    SchemaValidator["Schema Validator\nsrc/plugins/schema-validator.ts"]
    Registry["Plugin Registry\nsrc/plugins/registry.ts"]
    PluginRuntime["Plugin Runtime\nsrc/plugins/runtime/"]

    subgraph PluginRegistrations["What a Plugin Can Register"]
        HookReg["Hooks\n(lifecycle events)"]
        ChannelReg["Channels\n(new messaging platforms)"]
        ToolReg["Agent Tools\n(custom tool factories)"]
        CLIReg["CLI Commands\n(subcommands)"]
        ServiceReg["Services\n(background tasks)"]
        ProviderReg["AI Providers\n(model backends)"]
        HTTPReg["HTTP Routes\n(gateway endpoints)"]
    end

    PluginConfig --> Discovery
    BundledSources --> Discovery
    Discovery --> Loader
    Loader --> SchemaValidator
    SchemaValidator --> Registry
    Registry --> PluginRuntime
    Registry --> PluginRegistrations
```

## Plugin API Surface

```mermaid
classDiagram
    class OpenClawPluginApi {
        +registerHook(name, handler, opts)
        +registerChannel(registration)
        +registerTool(factory, names, opts)
        +registerCli(registrar, commands, opts)
        +registerService(service)
        +registerProvider(provider)
        +registerHttpRoute(params)
        +logger: PluginLogger
        +config: OpenClawConfig
    }
    class PluginHookRegistration {
        +name: PluginHookName
        +handler: PluginHookHandlerMap[name]
        +options?: PluginHookOptions
    }
    class PluginToolRegistration {
        +pluginId: string
        +factory: OpenClawPluginToolFactory
        +names: string[]
        +optional: boolean
        +source: string
    }
    class PluginCliRegistration {
        +pluginId: string
        +register: OpenClawPluginCliRegistrar
        +commands: string[]
        +source: string
    }

    OpenClawPluginApi --> PluginHookRegistration
    OpenClawPluginApi --> PluginToolRegistration
    OpenClawPluginApi --> PluginCliRegistration
```

## Plugin Lifecycle

```mermaid
stateDiagram-v2
    [*] --> Discovered : config / bundled sources
    Discovered --> Loading : loader.loadPlugin()
    Loading --> Validating : schema check
    Validating --> Registered : registry.registerPlugin()
    Registered --> Active : gateway start
    Active --> Disabled : toggle-config / uninstall
    Disabled --> Active : re-enable
    Active --> [*] : gateway shutdown
```

## Hook System

```mermaid
flowchart LR
    HookTrigger["Hook trigger\n(e.g. before-agent-run)"]
    HookRunner["Hook Runner\nsrc/plugins/hook-runner-global.ts"]
    InternalHooks["Internal Hooks\nsrc/hooks/internal-hooks.ts"]
    PluginHooks["Plugin Hooks\n(registered via registerHook)"]
    FireForget["fire-and-forget\nsrc/hooks/fire-and-forget.ts"]

    HookTrigger --> HookRunner
    HookRunner --> InternalHooks
    HookRunner --> PluginHooks
    PluginHooks --> FireForget
```

## Key Files

| File | Role |
|------|------|
| `src/plugins/discovery.ts` | Find installed plugins on disk |
| `src/plugins/loader.ts` | Dynamic import and initialise plugin module |
| `src/plugins/schema-validator.ts` | Validate plugin manifest schema |
| `src/plugins/registry.ts` | Central registry for all plugin registrations |
| `src/plugins/runtime/` | Runtime state for active plugins |
| `src/plugins/hook-runner-global.ts` | Execute hooks across all plugins |
| `src/plugins/commands.ts` | Register plugin-contributed CLI commands |
| `src/plugins/tools.ts` | Register plugin-contributed agent tools |
| `src/plugins/services.ts` | Manage plugin background services |
| `src/plugins/providers.ts` | Register plugin AI provider backends |
| `src/plugins/update.ts` | Plugin update / install flow |
| `src/plugins/uninstall.ts` | Plugin removal flow |
