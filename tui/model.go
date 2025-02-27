package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	interfaces map[string]string
	keys       []string
	cursor     int
	choices    map[string]string
	outputDir  string
	done       bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.done = true
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.keys)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			key := m.keys[m.cursor]
			fmt.Printf("\nEnter custom path for %s (default: %s): ", key, m.outputDir)
			var customPath string
			fmt.Scanln(&customPath)

			if strings.TrimSpace(customPath) == "" {
				customPath = m.outputDir
			}
			m.choices[key] = customPath
		case " ":
			key := m.keys[m.cursor]
			m.choices[key] = m.outputDir
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.done {
		return "Generating mocks...\n"
	}

	s := "Select interfaces to generate mocks for:\n\n"

	for i, key := range m.keys {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s -> %s\n", cursor, key, m.choices[key])
	}

	s += "\n[Enter] Customize path | [Space] Accept default | [Q] Quit\n"
	return s
}

func RunInterfaceSelector(interfaces map[string]string, outputDir string) (map[string]string, error) {
	keys := make([]string, 0, len(interfaces))
	choices := make(map[string]string)

	for k := range interfaces {
		keys = append(keys, k)
		choices[k] = outputDir
	}

	m := model{interfaces, keys, 0, choices, outputDir, false}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return nil, err
	}

	return m.choices, nil
}
