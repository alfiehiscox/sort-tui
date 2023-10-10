package main

import (
	"fmt"
	"os"

	"github.com/alfiehiscox/sort-tui/sorter"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {

	main := MainModel{
		Sorters:              sorter.Sorters,
		SelectedSorter:       0,
		FocusedSorter:        0,
		CodeView:             false,
		VisualisationRunning: false,
	}

	p := tea.NewProgram(main, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}

// MAIN UI Element
const (
	COLUMN_NUM        = 6 // How many vertical columns are there
	MARGIN_VERTICAL   = 0
	MARGIN_HORIZONTAL = 1
)

var ()

type MainModel struct {
	Width, Height        int
	ColumnWidth          int
	Sorters              []sorter.Sorter
	SelectedSorter       int // Actually being rendered
	FocusedSorter        int // Highlighted Over
	CodeView             bool
	VisualisationRunning bool
}

func (m MainModel) Init() tea.Cmd { return nil }

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width-MARGIN_HORIZONTAL, msg.Height-MARGIN_VERTICAL
		m.ColumnWidth = m.Width / COLUMN_NUM
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.SelectedSorter = m.FocusedSorter
		case "down", "j":
			if m.FocusedSorter < len(m.Sorters)-1 {
				m.FocusedSorter++
			}
		case "up", "k":
			if m.FocusedSorter > 0 {
				m.FocusedSorter--
			}
		case "right":
			m.CodeView = true
		case "left":
			m.CodeView = false
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	// Overall Styles
	subtle := lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	selected := lipgloss.Color("#55628F")

	base := lipgloss.NewStyle().
		Width(m.Width-1).
		Height(m.Height-2).
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(subtle)

	// Navigation Section (1 Column wide)
	navItem := lipgloss.NewStyle().
		Height(1).
		Width(m.ColumnWidth - 2)
	navTitle := navItem.Copy().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(subtle).
		Render("sort algorithm")

	var sorters string
	for i, sorter := range m.Sorters {
		if i == m.FocusedSorter {
			sorters += navItem.Copy().Background(selected).Render(sorter.Name()) + "\n"
		} else {
			sorters += navItem.Copy().Render(sorter.Name()) + "\n"
		}
	}

	nav := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		Height(base.GetHeight()).
		Width(m.ColumnWidth).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(
			lipgloss.Top,
			navTitle,
			sorters,
		))

	// Info/Code Section (3 Column wide)
	infoTitle := lipgloss.NewStyle().
		Height(1).
		Width(m.ColumnWidth*3-2).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(subtle).
		Render(m.Sorters[m.SelectedSorter].Name())

	detailsTab := lipgloss.NewStyle().
		Height(1).
		Width((m.ColumnWidth*3-2)/2).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(subtle).
		Align(lipgloss.Center)

	codeTab := lipgloss.NewStyle().
		Height(1).
		Width((m.ColumnWidth*3-2)/2).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(subtle).
		Align(lipgloss.Center)

	var tabs string
	if m.CodeView {
		tabs = lipgloss.JoinHorizontal(
			lipgloss.Left,
			detailsTab.Render("details"),
			codeTab.Background(selected).Render("code"),
		)
	} else {
		tabs = lipgloss.JoinHorizontal(
			lipgloss.Left,
			detailsTab.Background(selected).Render("details"),
			codeTab.Render("code"),
		)
	}

	info := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		Height(base.GetHeight()).
		Width(m.ColumnWidth*3).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(lipgloss.Top, infoTitle, tabs))

	// Visualisation Section
	vis := lipgloss.NewStyle().
		Height(base.GetHeight()).
		Width(m.ColumnWidth*2).
		Padding(0, 1).
		Render("Vis Section")

	return base.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		nav, info, vis,
	))
}
