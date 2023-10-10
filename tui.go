package main

import (
	"bytes"
	"fmt"

	"github.com/alfiehiscox/sort-tui/sorter"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const SIZE = 20

type Model struct {
	items []sorter.Item
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sorter.UpdateMsg:
		m.items = msg
	case sorter.FinishMsg:
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	var buf bytes.Buffer
	for _, item := range m.items {
		style := lipgloss.NewStyle()
		value := fmt.Sprintf("%d", item.Value)
		buf.WriteString(value)
		for i := len(value); i < 4; i++ {
			buf.WriteString(" ")
		}
		if item.Focused {
			bar := style.Foreground(lipgloss.Color("#359c32")).Render(item.Bar)
			buf.WriteString(bar + "\n")
		} else {
			bar := style.Foreground(lipgloss.Color("1")).Render(item.Bar)
			buf.WriteString(bar + "\n")
		}
	}
	return buf.String()
}
