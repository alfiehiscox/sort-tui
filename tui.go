package main

import (
	"bytes"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	items []Item
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SwapMsg:
		m.items[msg.IndexA], m.items[msg.IndexB] = m.items[msg.IndexB], m.items[msg.IndexA]
	case FinishMsg:
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
		bar := style.Foreground(item.Color).Render(item.Bar)
		buf.WriteString(bar + "\n")
	}
	return buf.String()
}
