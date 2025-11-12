package highlighter

import "github.com/charmbracelet/lipgloss"

var (
	KeywordStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#C678DD"))   // purple
	NamespaceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#56B6C2"))   // cyan
	FuncStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#0c06abff")) // blue
	TypeStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#E5C07B"))   // yellow
	StringStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#98C379"))   // green
	DefaultStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#47beeaff")) // gray
)

func styleForToken(t SemanticToken) lipgloss.Style {
	switch t.Type {
	case "keyword":
		return KeywordStyle
	case "namespace":
		return NamespaceStyle
	case "function":
		return FuncStyle
	case "type":
		return TypeStyle
	case "string":
		return StringStyle
	default:
		return DefaultStyle
	}
}
