# Authentication Profiles

How OpenClaw manages AI provider credentials — API keys, OAuth tokens, and multi-profile ordering.

```mermaid
flowchart TD
    AuthConfig["Auth config\nconfig.auth"]
    ProfileStore["Auth Profile Store\nsrc/agents/auth-profiles/store.ts"]
    Profiles["Profile Registry\nsrc/agents/auth-profiles/profiles.ts"]
    OrderResolver["Profile Order Resolver\nsrc/agents/auth-profiles/order.ts"]
    CredentialState["Credential State\nsrc/agents/auth-profiles/credential-state.ts"]
    OAuthProvider["OAuth Flow\nsrc/agents/auth-profiles/oauth.ts"]
    APIKeyStore["API Key Credential\n(stored in auth store)"]
    TokenCredential["Token Credential\n(short-lived token)"]
    ExternalCLISync["External CLI Sync\nsrc/agents/auth-profiles/external-cli-sync.ts"]
    Display["Display Helpers\nsrc/agents/auth-profiles/display.ts"]
    Doctor["Auth Doctor\nsrc/agents/auth-profiles/doctor.ts"]
    Usage["Usage Statistics\nsrc/agents/auth-profiles/usage.ts"]
    Repair["Repair Utilities\nsrc/agents/auth-profiles/repair.ts"]

    AuthConfig --> ProfileStore
    ProfileStore --> Profiles
    Profiles --> OrderResolver
    OrderResolver --> CredentialState
    CredentialState --> OAuthProvider
    CredentialState --> APIKeyStore
    CredentialState --> TokenCredential
    OAuthProvider --> ExternalCLISync
    Profiles --> Display
    Profiles --> Doctor
    Profiles --> Usage
    Profiles --> Repair
```

## Credential Types

```mermaid
classDiagram
    class AuthProfileCredential {
        <<union>>
        ApiKeyCredential
        OAuthCredential
        TokenCredential
    }
    class ApiKeyCredential {
        +type: "api-key"
        +apiKey: string
        +createdAt: number
    }
    class OAuthCredential {
        +type: "oauth"
        +accessToken: string
        +refreshToken?: string
        +expiresAt?: number
        +scopes: string[]
    }
    class TokenCredential {
        +type: "token"
        +token: string
        +expiresAt: number
    }

    AuthProfileCredential <|-- ApiKeyCredential
    AuthProfileCredential <|-- OAuthCredential
    AuthProfileCredential <|-- TokenCredential
```

## Profile Resolution Order

```mermaid
flowchart LR
    Start["resolveAuthProfileOrder()"]
    EligibilityCheck["resolveAuthProfileEligibility()\ncheck cooldown, expiry, health"]
    OrderSort["sort by priority + last-used"]
    CooldownCheck["isProfileInCooldown()"]
    HealthCheck["markAuthProfileGood/Failure()"]
    SelectedProfile["selected profile"]

    Start --> EligibilityCheck
    EligibilityCheck --> CooldownCheck
    CooldownCheck --> OrderSort
    OrderSort --> HealthCheck
    HealthCheck --> SelectedProfile
```

## Cooldown and Failure Tracking

```mermaid
flowchart TD
    AgentRun["Agent Run"]
    Success["Success"] --> MarkGood["markAuthProfileGood()"]
    Failure["API Error / Rate Limit"] --> MarkFail["markAuthProfileFailure()"]
    MarkFail --> Cooldown["markAuthProfileCooldown()\n(exponential backoff)"]
    Cooldown --> CooldownExpiry["getSoonestCooldownExpiry()"]
    CooldownExpiry --> NextProfile["try next eligible profile"]

    AgentRun --> Success
    AgentRun --> Failure
```

## OAuth Profile Session Override

```mermaid
sequenceDiagram
    participant CLI
    participant ProfileStore
    participant OAuth as oauth.ts
    participant ExternalCLI as External CLI (claude/codex)

    CLI->>ProfileStore: loadAuthProfileStore()
    ProfileStore-->>CLI: AuthProfileStore
    CLI->>OAuth: resolveApiKeyForProfile(profile)
    OAuth->>OAuth: check token expiry
    alt token expired
        OAuth->>ExternalCLI: refresh token (external-cli-sync.ts)
        ExternalCLI-->>OAuth: fresh token
        OAuth->>ProfileStore: upsertAuthProfile(updated)
    end
    OAuth-->>CLI: valid credential
```

## Key Files

| File | Role |
|------|------|
| `src/agents/auth-profiles/store.ts` | Load/save auth profile store (JSON on disk) |
| `src/agents/auth-profiles/profiles.ts` | Upsert, dedupe, list profiles |
| `src/agents/auth-profiles/order.ts` | Eligibility check and ordering logic |
| `src/agents/auth-profiles/credential-state.ts` | Token expiry, state classification |
| `src/agents/auth-profiles/oauth.ts` | OAuth token resolution and refresh |
| `src/agents/auth-profiles/external-cli-sync.ts` | Sync tokens from external CLI tools |
| `src/agents/auth-profiles/types.ts` | TypeScript credential type definitions |
| `src/agents/auth-profiles/usage.ts` | Track per-profile usage statistics |
| `src/agents/auth-profiles/repair.ts` | Fix mismatched profile IDs |
| `src/agents/auth-profiles/display.ts` | Format profiles for terminal display |
| `src/agents/auth-profiles/doctor.ts` | Diagnose auth configuration issues |
| `src/agents/auth-profiles/paths.ts` | Resolve auth store file paths |
