package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TUIOptions holds options for the TUI.
type TUIOptions struct {
	URL          string
	Token        string
	Password     string
	SessionKey   string
	AgentID      string
	Deliver      string
	Thinking     string
	Message      string
	TimeoutMs    int
	HistoryLimit int
	GatewayPort  int
}

// TUIState represents the current state of the TUI.
type TUIState string

const (
	TUIStateConnecting   TUIState = "connecting"
	TUIStateConnected    TUIState = "connected"
	TUIStateWaiting      TUIState = "waiting"
	TUIStateDisconnected TUIState = "disconnected"
	TUIStateError        TUIState = "error"
)

// Model is the BubbleTea model for the TUI.
type Model struct {
	opts           TUIOptions
	state          TUIState
	chatLog        *ChatLog
	viewport       viewport.Model
	textarea       textarea.Model
	client         *GatewayChatClient
	width          int
	height         int
	statusMsg      string
	err            error
	history        []string
	historyPos     int
	currentSession string
	currentAgent   string
	thinking       bool
	streamBuf      strings.Builder
}

// Styles
var (
	statusBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#1C1C1C")).
			Foreground(lipgloss.Color("#AAAAAA")).
			Padding(0, 1)

	connectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00C853"))

	disconnectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF1744"))

	waitingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFAB00"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#484848")).
			Faint(true)

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#333333"))
)

// tickMsg is used for periodic updates.
type tickMsg time.Time

// gatewayEventMsg wraps a gateway event for BubbleTea.
type gatewayEventMsg struct {
	event GatewayEvent
}

// connectedMsg signals successful connection.
type connectedMsg struct{}

// disconnectedMsg signals disconnection.
type disconnectedMsg struct {
	err error
}

// NewModel creates a new TUI model.
func NewModel(opts TUIOptions) *Model {
	if opts.HistoryLimit <= 0 {
		opts.HistoryLimit = 200
	}
	if opts.GatewayPort <= 0 {
		opts.GatewayPort = 18789
	}

	ta := textarea.New()
	ta.Placeholder = "Type a message... (/help for commands)"
	ta.Focus()
	ta.SetWidth(80)
	ta.SetHeight(3)
	ta.ShowLineNumbers = false
	ta.CharLimit = 4000

	vp := viewport.New(80, 20)

	connOpts := ResolveGatewayConnection(opts.URL, opts.Token, opts.Password, opts.GatewayPort)
	client := NewGatewayChatClient(connOpts)

	m := &Model{
		opts:           opts,
		state:          TUIStateConnecting,
		chatLog:        NewChatLog(opts.HistoryLimit, 80),
		viewport:       vp,
		textarea:       ta,
		client:         client,
		statusMsg:      "Connecting...",
		historyPos:     -1,
		currentSession: opts.SessionKey,
		currentAgent:   opts.AgentID,
	}
	return m
}

// Init initializes the TUI model.
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		m.connect(),
		tick(),
	)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) connect() tea.Cmd {
	return func() tea.Msg {
		m.client.OnEvent = func(event GatewayEvent) {}
		m.client.OnConnected = func() {}
		m.client.OnDisconnected = func(err error) {}
		if err := m.client.Connect(); err != nil {
			return disconnectedMsg{err: err}
		}
		return connectedMsg{}
	}
}

// Update handles messages and updates the model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		cmds  []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - 8
		m.textarea.SetWidth(msg.Width - 4)
		m.chatLog.width = msg.Width
		m.updateViewport()

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			if m.state == TUIStateWaiting {
				m.client.AbortChat(m.currentSession)
				m.state = TUIStateConnected
				m.statusMsg = "Aborted"
			}
		case tea.KeyEnter:
			if msg.Alt {
				// Alt+Enter: newline in textarea
				m.textarea, tiCmd = m.textarea.Update(msg)
				cmds = append(cmds, tiCmd)
			} else {
				input := strings.TrimSpace(m.textarea.Value())
				if input != "" {
					cmds = append(cmds, m.handleInput(input))
					m.textarea.Reset()
				}
			}
			return m, tea.Batch(cmds...)
		case tea.KeyUp:
			// Navigate history
			if m.historyPos < len(m.history)-1 {
				m.historyPos++
				m.textarea.SetValue(m.history[len(m.history)-1-m.historyPos])
			}
		case tea.KeyDown:
			// Navigate history
			if m.historyPos > 0 {
				m.historyPos--
				m.textarea.SetValue(m.history[len(m.history)-1-m.historyPos])
			} else if m.historyPos == 0 {
				m.historyPos = -1
				m.textarea.SetValue("")
			}
		}

	case tickMsg:
		cmds = append(cmds, tick())

	case connectedMsg:
		m.state = TUIStateConnected
		m.statusMsg = "Connected"
		m.chatLog.Add(ChatMessage{
			Role:    MessageRoleSystem,
			Content: "Connected to gateway",
		})
		m.updateViewport()
		if m.opts.Message != "" {
			cmds = append(cmds, m.handleInput(m.opts.Message))
		}

	case disconnectedMsg:
		m.state = TUIStateDisconnected
		if msg.err != nil {
			m.statusMsg = "Disconnected: " + msg.err.Error()
		} else {
			m.statusMsg = "Disconnected"
		}
		m.chatLog.Add(ChatMessage{
			Role:    MessageRoleSystem,
			Content: "Disconnected from gateway",
		})
		m.updateViewport()

	case gatewayEventMsg:
		cmds = append(cmds, m.handleGatewayEvent(msg.event))
	}

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
	cmds = append(cmds, tiCmd, vpCmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) handleInput(input string) tea.Cmd {
	// Handle slash commands
	if strings.HasPrefix(input, "/") {
		return m.handleCommand(input)
	}

	// Add to history
	m.history = append(m.history, input)
	m.historyPos = -1

	// Add user message to chat log
	m.chatLog.Add(ChatMessage{
		Role:    MessageRoleUser,
		Content: input,
	})
	m.updateViewport()
	m.state = TUIStateWaiting
	m.statusMsg = "Waiting for response..."

	// Send to gateway
	return func() tea.Msg {
		err := m.client.SendChat(ChatSendOptions{
			Message:    input,
			SessionKey: m.currentSession,
			AgentID:    m.currentAgent,
			Thinking:   m.opts.Thinking,
			Deliver:    m.opts.Deliver,
			TimeoutMs:  m.opts.TimeoutMs,
		})
		if err != nil {
			return gatewayEventMsg{event: GatewayEvent{
				Type:  GatewayEventTypeError,
				Error: err.Error(),
			}}
		}
		return nil
	}
}

func (m *Model) handleCommand(input string) tea.Cmd {
	parts := strings.Fields(input)
	cmd := parts[0]
	switch cmd {
	case "/help":
		m.chatLog.Add(ChatMessage{
			Role: MessageRoleSystem,
			Content: `Available commands:
  /help          - Show this help
  /sessions      - List sessions
  /agents        - List agents
  /models        - List models
  /clear         - Clear chat log
  /quit, /exit   - Exit TUI
  Ctrl+C         - Exit
  Esc            - Abort current request
  Up/Down        - Navigate history`,
		})
		m.updateViewport()
	case "/clear":
		m.chatLog.Clear()
		m.updateViewport()
	case "/quit", "/exit":
		return tea.Quit
	case "/sessions":
		m.client.ListSessions()
	case "/agents":
		m.client.ListAgents()
	case "/models":
		m.client.ListModels()
	}
	return nil
}

func (m *Model) handleGatewayEvent(event GatewayEvent) tea.Cmd {
	switch event.Type {
	case GatewayEventTypeStream:
		m.streamBuf.WriteString(event.Text)
		// Update last assistant message if streaming
		msgs := m.chatLog.Messages()
		if len(msgs) > 0 && msgs[len(msgs)-1].Role == MessageRoleAssistant && msgs[len(msgs)-1].IsStreaming {
			m.chatLog.messages[len(m.chatLog.messages)-1].Content = m.streamBuf.String()
		} else {
			m.streamBuf.Reset()
			m.streamBuf.WriteString(event.Text)
			m.chatLog.Add(ChatMessage{
				Role:        MessageRoleAssistant,
				Content:     m.streamBuf.String(),
				IsStreaming: true,
			})
		}
		m.updateViewport()

	case GatewayEventTypeChat:
		// Final message
		if text, ok := event.Payload["text"].(string); ok {
			msgs := m.chatLog.Messages()
			if len(msgs) > 0 && msgs[len(msgs)-1].IsStreaming {
				m.chatLog.messages[len(m.chatLog.messages)-1].Content = text
				m.chatLog.messages[len(m.chatLog.messages)-1].IsStreaming = false
			} else {
				m.chatLog.Add(ChatMessage{
					Role:    MessageRoleAssistant,
					Content: text,
				})
			}
		}
		m.streamBuf.Reset()
		m.state = TUIStateConnected
		m.statusMsg = "Connected"
		m.updateViewport()

	case GatewayEventTypeDone:
		m.state = TUIStateConnected
		m.statusMsg = "Connected"

	case GatewayEventTypeToolStart:
		toolName := ""
		if name, ok := event.Payload["name"].(string); ok {
			toolName = name
		}
		m.chatLog.Add(ChatMessage{
			Role:     MessageRoleTool,
			Content:  "Starting...",
			ToolName: toolName,
		})
		m.updateViewport()

	case GatewayEventTypeToolEnd:
		toolName := ""
		if name, ok := event.Payload["name"].(string); ok {
			toolName = name
		}
		result := "Done"
		if r, ok := event.Payload["result"].(string); ok {
			result = r
		}
		m.chatLog.Add(ChatMessage{
			Role:     MessageRoleTool,
			Content:  result,
			ToolName: toolName,
		})
		m.updateViewport()

	case GatewayEventTypeError:
		m.chatLog.Add(ChatMessage{
			Role:    MessageRoleSystem,
			Content: "Error: " + event.Error,
		})
		m.state = TUIStateConnected
		m.statusMsg = "Error"
		m.updateViewport()
	}
	return nil
}

func (m *Model) updateViewport() {
	m.viewport.SetContent(m.chatLog.Render())
	m.viewport.GotoBottom()
}

// View renders the TUI.
func (m *Model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Status bar
	stateLabel := m.renderStateLabel()
	agent := m.currentAgent
	if agent == "" {
		agent = "main"
	}
	session := m.currentSession
	if session == "" {
		session = "default"
	}
	statusBar := statusBarStyle.Width(m.width).Render(
		fmt.Sprintf(" %s  agent: %s  session: %s  %s",
			stateLabel, agent, session, m.statusMsg),
	)

	// Chat log viewport
	chatView := borderStyle.
		Width(m.width - 2).
		Height(m.height - 8).
		Render(m.viewport.View())

	// Input area
	inputView := borderStyle.
		Width(m.width - 2).
		Render(m.textarea.View())

	// Help bar
	helpBar := helpStyle.Render(
		"Enter: send  Alt+Enter: newline  ↑↓: history  Esc: abort  Ctrl+C: quit  /help: commands",
	)

	return lipgloss.JoinVertical(lipgloss.Left,
		statusBar,
		chatView,
		inputView,
		helpBar,
	)
}

func (m *Model) renderStateLabel() string {
	switch m.state {
	case TUIStateConnected:
		return connectedStyle.Render("●")
	case TUIStateWaiting:
		return waitingStyle.Render("◌")
	case TUIStateConnecting:
		return waitingStyle.Render("◌")
	case TUIStateDisconnected:
		return disconnectedStyle.Render("○")
	case TUIStateError:
		return disconnectedStyle.Render("✗")
	default:
		return "?"
	}
}

// RunTUI runs the TUI application.
func RunTUI(opts TUIOptions) error {
	m := NewModel(opts)
	p := tea.NewProgram(m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	_, err := p.Run()
	return err
}
