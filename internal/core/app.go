package core

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focusedScreen int

const (
	fileExplorerScreen focusedScreen = iota
	textEditorScreen
	terminalScreen
)

type AppModel struct {
	fileExplorer fileExplorerModel
	textEditor   textEditorModel
	focused      focusedScreen
	width        int
	height       int
	someText     string
}

func NewApp() AppModel {
	return AppModel{
		fileExplorer: newFileExplorer(),
		textEditor:   newTextEditor(),
		focused:      fileExplorerScreen,
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(m.fileExplorer.Init(), m.textEditor.Init())
}

func (m AppModel) View() string {
	m.fileExplorer.SetFocused(m.focused == fileExplorerScreen)
	m.textEditor.SetFocused(m.focused == textEditorScreen)
	return lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		m.fileExplorer.View(),
		m.textEditor.View(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - 4
		m.height = msg.Height - 2
		m.fileExplorer.SetSize(m.width/4, m.height)
		m.textEditor.SetSize(m.width-(m.width/4), m.height)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEsc.String():
			// m. = true
			return m, tea.Quit
		case "alt+E":
			if m.focused != fileExplorerScreen {
				m.fileExplorer.SetSize(m.width/4, m.height)
				m.focused = fileExplorerScreen
			} else {
				m.focused = textEditorScreen
			}
		case "ctrl+b":
			if m.fileExplorer.width > 0 {
				m.fileExplorer.SetSize(0, m.height)
				if m.focused == fileExplorerScreen {
					m.focused = fileExplorerScreen
				}
			} else {
				m.fileExplorer.SetSize(m.width/4, m.height)
			}
			m.focused = textEditorScreen
		}

	}
	var cmd tea.Cmd
	switch m.focused {
	case fileExplorerScreen:
		m.fileExplorer, cmd = m.fileExplorer.Update(msg)
	case textEditorScreen:
		m.textEditor, cmd = m.textEditor.Update(msg)
	}

	return m, cmd
}

func (m AppModel) UpdateFocused() {
	screen := m.focused
	m.fileExplorer.SetFocused(screen == fileExplorerScreen)
	m.textEditor.SetFocused(screen == textEditorScreen)
}
