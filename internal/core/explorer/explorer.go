package explorer

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FileItem struct {
	name  string
	path  string
	isDir bool
}

func (i FileItem) Title() string {
	icon := ""
	if i.isDir {
		// isOpen := "📂"
		icon = ""
	}

	return fmt.Sprintf("%s %s", icon, i.name)
}

type fileExplorerModel struct {
	filepicker   filepicker.Model
	selectedFile string
	quitting     bool
	err          error
	focused      bool
	width        int
	height       int
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func newFileExplorer() fileExplorerModel {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".mod", ".sum", ".go", ".txt", ".md"}
	fp.CurrentDirectory, _ = os.UserHomeDir()

	return fileExplorerModel{
		filepicker: fp,
	}
}

func (m fileExplorerModel) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m fileExplorerModel) Update(msg tea.Msg) (fileExplorerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		m.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		m.err = errors.New(path + " is not valid.")
		m.selectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m fileExplorerModel) View() string {
	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#5A5A5A")).
		Padding(1, 0, 0).
		Width(m.width).
		Height(m.height)

	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.selectedFile == "" {
		s.WriteString("Pick a file:")
	} else {
		s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
	}
	s.WriteString("\n\n" + m.filepicker.View() + "\n")

	if m.focused {
		borderStyle = borderStyle.BorderForeground(lipgloss.Color("#00BFFF"))
	}

	return borderStyle.Render(s.String())
}

func (m *fileExplorerModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func (m *fileExplorerModel) SetFocused(f bool) {
	m.focused = f
}
