# Process Supervision

How OpenClaw spawns, tracks, and cancels agent child processes through the supervisor layer.

```mermaid
flowchart TD
    AgentCLIRunner["Agent CLI Runner\nsrc/agents/cli-runner.ts"]
    Supervisor["Process Supervisor\nsrc/process/supervisor/supervisor.ts"]
    Registry["Run Registry\nsrc/process/supervisor/registry.ts"]

    subgraph Adapters["Process Adapters"]
        ChildAdapter["Child Process\nsrc/process/supervisor/adapters/child.ts"]
        PtyAdapter["PTY\nsrc/process/supervisor/adapters/pty.ts"]
        EnvAdapter["Env Override\nsrc/process/supervisor/adapters/env.ts"]
    end

    Lanes["Process Lanes\nsrc/process/lanes.ts"]
    CommandQueue["Command Queue\nsrc/process/command-queue.ts"]
    KillTree["Kill Tree\nsrc/process/kill-tree.ts"]
    Exec["Exec Helper\nsrc/process/exec.ts"]
    RestartRecovery["Restart Recovery\nsrc/process/restart-recovery.ts"]
    ChildProcessBridge["Child-Process Bridge\nsrc/process/child-process-bridge.ts"]
    SpawnUtils["Spawn Utilities\nsrc/process/spawn-utils.ts"]

    AgentCLIRunner --> Supervisor
    Supervisor --> Registry
    Supervisor --> Adapters
    Supervisor --> KillTree
    Adapters --> Lanes
    Lanes --> CommandQueue
    CommandQueue --> Exec
    Exec --> SpawnUtils
    RestartRecovery --> Supervisor
    ChildProcessBridge --> SpawnUtils
```

## Process Supervisor State Machine

```mermaid
stateDiagram-v2
    [*] --> starting : spawn(input)
    starting --> running : first output received
    running --> exiting : cancel() / timeout
    running --> exited : process exits 0
    exiting --> exited : process exits (any code)
    exited --> [*]
```

## Process Supervisor API

```mermaid
classDiagram
    class ProcessSupervisor {
        +spawn(input: SpawnInput) Promise~ManagedRun~
        +cancel(runId, reason)
        +cancelScope(scopeKey, reason)
        +getState(runId) RunRecord
        +listActive() RunRecord[]
    }
    class SpawnInput {
        +runId?: string
        +sessionId: string
        +backendId: string
        +scopeKey?: string
        +replaceExistingScope?: boolean
        +command: string
        +args: string[]
        +cwd?: string
        +env?: Record~string, string~
        +timeoutMs?: number
        +noOutputTimeoutMs?: number
        +usePty?: boolean
    }
    class RunRecord {
        +runId: string
        +sessionId: string
        +backendId: string
        +scopeKey?: string
        +state: RunState
        +startedAtMs: number
        +lastOutputAtMs: number
        +terminationReason?: TerminationReason
    }
    class ManagedRun {
        +runId: string
        +cancel(reason)
        +output AsyncIterable~string~
        +exitCode Promise~number~
    }

    ProcessSupervisor --> SpawnInput : accepts
    ProcessSupervisor --> ManagedRun : returns
    ProcessSupervisor --> RunRecord : tracks
```

## Lanes and Concurrency

```mermaid
graph LR
    Supervisor["Process Supervisor"]
    Supervisor --> LaneA["Lane: agent\n(AI runs, concurrency limited)"]
    Supervisor --> LaneB["Lane: bash\n(shell tool executions)"]
    Supervisor --> LaneC["Lane: default\n(misc processes)"]
    LaneA --> CommandQueue["Command Queue\n(per-scope FIFO)"]
    LaneB --> CommandQueue
    LaneC --> CommandQueue
```

## Termination Reasons

```mermaid
graph TD
    TermReason["TerminationReason"]
    TermReason --> ManualCancel["manual-cancel\n(user/system cancel)"]
    TermReason --> OverallTimeout["overall-timeout\n(wall-clock limit exceeded)"]
    TermReason --> NoOutputTimeout["no-output-timeout\n(stall watchdog)"]
    TermReason --> ScopeReplace["scope-replace\n(new run replaced this scope)"]
```

## Key Files

| File | Role |
|------|------|
| `src/process/supervisor/supervisor.ts` | Core supervisor: spawn, cancel, track |
| `src/process/supervisor/registry.ts` | In-memory run registry |
| `src/process/supervisor/adapters/child.ts` | Node `child_process.spawn` adapter |
| `src/process/supervisor/adapters/pty.ts` | PTY (pseudo-terminal) adapter |
| `src/process/supervisor/adapters/env.ts` | Environment variable injection |
| `src/process/lanes.ts` | Process lane definitions and concurrency |
| `src/process/command-queue.ts` | Per-scope sequential command queue |
| `src/process/kill-tree.ts` | Recursively kill process tree |
| `src/process/exec.ts` | Low-level exec helper |
| `src/process/spawn-utils.ts` | Cross-platform spawn utilities |
| `src/process/restart-recovery.ts` | Supervisor restart/recovery logic |
| `src/process/child-process-bridge.ts` | IPC bridge for child processes |
