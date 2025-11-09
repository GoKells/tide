package core

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type textEditorModel struct {
	textarea textarea.Model
	err      error
	focused  bool
	width    int
	height   int
}

type errMsg error

func newTextEditor() textEditorModel {
	ti := textarea.New()
	ti.Placeholder = "Text Editor"
	ti.Focus()
	ti.CharLimit = 0
	ti.ShowLineNumbers = true
	ti.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("#1E1E1E"))
	ti.Prompt = ""

	return textEditorModel{
		textarea: ti,
		err:      nil,
	}
}

func (m textEditorModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m textEditorModel) Update(msg tea.Msg) (textEditorModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m textEditorModel) View() string {
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#5A5A5A")).
		Padding(0).
		Width(m.width).
		Height(m.height)

	if m.focused {
		borderStyle = borderStyle.BorderForeground(lipgloss.Color("#00BFFF"))
	}

	m.textarea.SetWidth(m.width - 2)
	m.textarea.SetHeight(m.height - 2)

	return borderStyle.Render(m.textarea.View())
}

func (m *textEditorModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *textEditorModel) SetFocused(f bool) {
	m.focused = f
}
