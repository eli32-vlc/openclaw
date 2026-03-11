package agents

// ProviderType identifies an AI provider.
type ProviderType string

const (
	ProviderOpenAI    ProviderType = "openai"
	ProviderAnthropic ProviderType = "anthropic"
	ProviderGemini    ProviderType = "gemini"
	ProviderCopilot   ProviderType = "copilot"
	ProviderMistral   ProviderType = "mistral"
	ProviderCohere    ProviderType = "cohere"
	ProviderGroq      ProviderType = "groq"
	ProviderOllama    ProviderType = "ollama"
	ProviderDeepSeek  ProviderType = "deepseek"
	ProviderXAI       ProviderType = "xai"
)

// AuthProfile holds credentials for a provider.
type AuthProfile struct {
	ID       string       `json:"id"`
	Provider ProviderType `json:"provider"`
	Mode     string       `json:"mode"`
	APIKey   string       `json:"apiKey,omitempty"`
	Email    string       `json:"email,omitempty"`
}

// Provider defines the interface for AI providers.
type Provider interface {
	ID() ProviderType
	Name() string
	DefaultModel() string
	Models() []string
}

// ProviderMeta holds metadata for a provider.
type ProviderMeta struct {
	id           ProviderType
	name         string
	defaultModel string
	models       []string
}

// ID returns the provider ID.
func (p *ProviderMeta) ID() ProviderType { return p.id }

// Name returns the provider name.
func (p *ProviderMeta) Name() string { return p.name }

// DefaultModel returns the default model.
func (p *ProviderMeta) DefaultModel() string { return p.defaultModel }

// Models returns available models.
func (p *ProviderMeta) Models() []string { return p.models }

// KnownProviders is the list of known AI providers.
var KnownProviders = []*ProviderMeta{
	{ProviderOpenAI, "OpenAI", "gpt-4o", []string{"gpt-4o", "gpt-4o-mini", "o1", "o3-mini"}},
	{ProviderAnthropic, "Anthropic", "claude-opus-4-5", []string{"claude-opus-4-5", "claude-sonnet-4-5", "claude-haiku-3-5"}},
	{ProviderGemini, "Google Gemini", "gemini-2.0-flash", []string{"gemini-2.0-flash", "gemini-2.0-pro"}},
	{ProviderCopilot, "GitHub Copilot", "gpt-4o", []string{"gpt-4o", "claude-opus-4-5"}},
	{ProviderMistral, "Mistral", "mistral-large", []string{"mistral-large", "mistral-small"}},
	{ProviderGroq, "Groq", "llama-3.3-70b-versatile", []string{"llama-3.3-70b-versatile"}},
	{ProviderOllama, "Ollama", "llama3.2", []string{"llama3.2", "mistral", "phi3"}},
}
