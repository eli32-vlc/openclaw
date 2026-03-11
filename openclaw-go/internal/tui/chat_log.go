package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// MessageRole identifies the role of a chat message.
type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleSystem    MessageRole = "system"
	MessageRoleTool      MessageRole = "tool"
)

// ChatMessage is a single message in the chat log.
type ChatMessage struct {
	Role        MessageRole
	Content     string
	Timestamp   time.Time
	ToolName    string
	IsStreaming bool
}

// ChatLog manages the list of chat messages.
type ChatLog struct {
	messages     []ChatMessage
	historyLimit int
	width        int
}

// NewChatLog creates a new chat log.
func NewChatLog(historyLimit, width int) *ChatLog {
	if historyLimit <= 0 {
		historyLimit = 200
	}
	return &ChatLog{historyLimit: historyLimit, width: width}
}

// Add adds a message to the log.
func (l *ChatLog) Add(msg ChatMessage) {
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}
	l.messages = append(l.messages, msg)
	if len(l.messages) > l.historyLimit {
		l.messages = l.messages[len(l.messages)-l.historyLimit:]
	}
}

// Messages returns all messages.
func (l *ChatLog) Messages() []ChatMessage {
	return l.messages
}

// Render renders the chat log as a string.
func (l *ChatLog) Render() string {
	userStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00B4D8")).
		Bold(true)
	assistantStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#90E0EF"))
	systemStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#757575")).
		Italic(true)
	toolStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFB703"))
	timestampStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#484848")).
		Faint(true)

	var sb strings.Builder
	for _, msg := range l.messages {
		ts := timestampStyle.Render(msg.Timestamp.Format("15:04:05"))
		switch msg.Role {
		case MessageRoleUser:
			prefix := userStyle.Render("You")
			sb.WriteString(fmt.Sprintf("%s %s: %s\n", ts, prefix, msg.Content))
		case MessageRoleAssistant:
			prefix := assistantStyle.Render("Assistant")
			sb.WriteString(fmt.Sprintf("%s %s: %s\n", ts, prefix, msg.Content))
		case MessageRoleSystem:
			sb.WriteString(fmt.Sprintf("%s %s\n", ts, systemStyle.Render(msg.Content)))
		case MessageRoleTool:
			name := msg.ToolName
			if name == "" {
				name = "tool"
			}
			prefix := toolStyle.Render("⚙ " + name)
			sb.WriteString(fmt.Sprintf("%s %s: %s\n", ts, prefix, msg.Content))
		}
	}
	return sb.String()
}

// Clear removes all messages.
func (l *ChatLog) Clear() {
	l.messages = nil
}

// Len returns the number of messages.
func (l *ChatLog) Len() int {
	return len(l.messages)
}
