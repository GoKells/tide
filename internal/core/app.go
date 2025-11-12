package core

import (
	"github.com/GoKells/tide/internal/core/editor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focusedScreen int

const (
	textEditorScreen focusedScreen = iota
)

type AppModel struct {
	// fileExplorer fileExplorerModel
	textEditor editor.EditorModel
	focused    focusedScreen
	width      int
	height     int
}

func NewApp() AppModel {
	return AppModel{
		// fileExplorer: newFileExplorer(),
		textEditor: editor.NewEditor(),
		focused:    textEditorScreen,
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(m.textEditor.Init())
}

func (m AppModel) View() string {
	m.textEditor.SetFocused(m.focused == textEditorScreen)
	return lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		m.textEditor.View(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width - 4
		m.height = msg.Height - 2
		// m.fileExplorer.SetSize(m.width/4, m.height)
		m.textEditor.SetSize(m.width-(m.width/4), m.height)
		return m, nil

	case tea.KeyMsg:
		var cmd tea.Cmd
		m.textEditor, cmd = m.textEditor.Update(msg)
		return m, cmd
	}

	return m, nil
}
