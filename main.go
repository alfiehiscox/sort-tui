package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alfiehiscox/sort-tui/sorter"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var p *tea.Program

func main() {

	main := MainModel{
		Sorters:              sorter.Sorters,
		SelectedSorter:       0,
		FocusedSorter:        0,
		VisualisationRunning: false,
		VisualisationRan:     false,
		// Items:                sorter.GetRandomItems(VISUALISATION_SIZE),
		Sub: make(chan []sorter.Item),
	}

	p = tea.NewProgram(main, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}

const (
	COLUMN_NUM        = 6 // How many vertical columns are there
	MARGIN_VERTICAL   = 0
	MARGIN_HORIZONTAL = 1
	// VISUALISATION_SIZE = 19
)

var ()

type MainModel struct {
	Width, Height         int
	ColumnWidth           int
	Sorters               []sorter.Sorter
	SelectedSorter        int // Actually being rendered
	FocusedSorter         int // Highlighted Over
	VisualisationRunning  bool
	VisualisationRan      bool
	VisualistationMaxElem int // The max elements to fit in the visualisation

	Items []sorter.Item // Item Visualiser
	Sub   chan []sorter.Item
}

func (m MainModel) Init() tea.Cmd { return nil }

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width-MARGIN_HORIZONTAL, msg.Height-MARGIN_VERTICAL
		if m.ColumnWidth != m.Width/COLUMN_NUM {
			m.ColumnWidth = m.Width / COLUMN_NUM
			maxCell := 2*m.ColumnWidth - 2 - 2
			m.VisualistationMaxElem = sorter.MaxArrayFromCells(maxCell)
			m.Items = sorter.GetRandomItems(m.VisualistationMaxElem)
		}
		return m, nil
	case sorter.UpdateMsg:
		m.Items = msg
		sorter := m.Sorters[m.SelectedSorter]
		return m, sorter.WaitForSort(m.Sub)
	case sorter.FinishMsg:
		m.VisualisationRunning = false
		m.VisualisationRan = true
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.VisualisationRunning {
				return m, nil
			}
			m.SelectedSorter = m.FocusedSorter
		case "down", "j":
			if m.VisualisationRunning {
				return m, nil
			}
			if m.FocusedSorter < len(m.Sorters)-1 {
				m.FocusedSorter++
			}
		case "up", "k":
			if m.VisualisationRunning {
				return m, nil
			}
			if m.FocusedSorter > 0 {
				m.FocusedSorter--
			}
		case "r":
			// Randomise
			if !m.VisualisationRunning {
				m.Items = sorter.GetRandomItems(m.VisualistationMaxElem)
				m.VisualisationRan = false
			}
		case " ":
			// Start Sorting Process
			if !m.VisualisationRunning && !m.VisualisationRan {
				m.VisualisationRunning = true
				sorter := m.Sorters[m.SelectedSorter]
				return m, tea.Batch(
					sorter.Sort(m.Items, m.Sub),
					sorter.WaitForSort(m.Sub),
				)
			}
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
		Render("Sort Algorithm")

	var sorters string
	for i, sorter := range m.Sorters {
		if i == m.FocusedSorter {
			sorters += navItem.Copy().
				Background(selected).
				Render(sorter.Name()) + "\n"
		} else {
			sorters += navItem.Copy().
				Render(sorter.Name()) + "\n"
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

	// Info Section (3 Column wide)
	info := lipgloss.NewStyle().Height(1).Width(m.ColumnWidth*3 - 2).Render("Information")
	infoTitle := lipgloss.NewStyle().
		Height(1).
		Width(m.ColumnWidth*3-2).
		Border(lipgloss.NormalBorder(), true, false, true, false).
		BorderForeground(subtle).
		Render(m.Sorters[m.SelectedSorter].Name())
	descriptionText := wordwrap.String(m.Sorters[m.SelectedSorter].Description(), m.ColumnWidth*3-2)
	description := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Render(descriptionText + "\n\n")
	complexity := "Complexity: \n"
	complexity += RenderComplexity(m.Sorters[m.SelectedSorter].Complexity(), m.ColumnWidth*3-2)
	main := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		Height(base.GetHeight()).
		Width(m.ColumnWidth*3).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(lipgloss.Top, info, infoTitle, description, complexity))

	// Visualisation Section (2 Column wide)
	visGap := lipgloss.NewStyle().Height(1).Width(m.ColumnWidth*2 - 2).Render()
	visTitle := lipgloss.NewStyle().
		Height(1).
		Width(m.ColumnWidth*2-2).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(subtle).
		Render("Visualisation")
	bar := lipgloss.NewStyle().
		Width(m.ColumnWidth*2 - 2).
		Align(lipgloss.Left)

	var xs bytes.Buffer
	xs.WriteByte('(')
	bars := []string{}
	for i, item := range m.Items {
		// Do some colouring thing here probably
		elem := ""
		if i == len(m.Items)-1 {
			elem = "%d)"
		} else {
			elem = "%d,"
		}
		if !item.Focused {
			xs.WriteString(fmt.Sprintf(elem, item.Value))
		} else {
			xs.WriteString(
				lipgloss.NewStyle().Foreground(selected).Render(
					fmt.Sprintf(elem, item.Value),
				),
			)
		}

		if item.Focused {
			bars = append(bars, bar.Copy().Foreground(selected).Render(item.Bar))
		} else {
			bars = append(bars, bar.Render(item.Bar))
		}
	}

	visSlice := lipgloss.NewStyle().
		Width(m.ColumnWidth*2 - 2).
		Align(lipgloss.Center).
		Render(xs.String())
	visBars := lipgloss.NewStyle().
		Width(m.ColumnWidth*2 - 2).
		PaddingLeft(2).
		Render(lipgloss.JoinVertical(lipgloss.Top, bars...))

	vis := lipgloss.NewStyle().
		Height(base.GetHeight()).
		Width(m.ColumnWidth*2).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(lipgloss.Top, visTitle, visGap, visSlice, visGap, visBars))

	return base.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		nav, main, vis,
	))
}

func RenderComplexity(c sorter.Complexity, width int) string {
	subtle := lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	selected := lipgloss.Color("#55628F")

	columnWidth := width / 4
	cell := lipgloss.NewStyle().
		Width(columnWidth).
		Height(1).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(subtle).Align(lipgloss.Left).
		PaddingLeft(1)
	time := cell.Copy().
		Border(lipgloss.NormalBorder(), true, false, true, false).
		Width(columnWidth * 3).Render("Time")
	space := cell.Copy().
		Border(lipgloss.NormalBorder(), true, false, true, false).
		Render("Space")
	timeBest := cell.Render("Best")
	timeAvg := cell.Render("Avg")
	timeWorst := cell.Render("Worst")
	spaceWorst := cell.Render("Worst")
	acTimeBest := cell.Background(selected).Render(c.TimeBest)
	acTimeAvg := cell.Render(c.TimeAvg)
	acTimeWorst := cell.Render(c.TimeWorst)
	acSpaceWorst := cell.Render(c.SpaceWorst)

	row1 := lipgloss.JoinHorizontal(lipgloss.Left, time, space)
	row2 := lipgloss.JoinHorizontal(lipgloss.Left, timeBest, timeAvg, timeWorst, spaceWorst)
	row3 := lipgloss.JoinHorizontal(lipgloss.Left, acTimeBest, acTimeAvg, acTimeWorst, acSpaceWorst)
	return lipgloss.JoinVertical(lipgloss.Top, row1, row2, row3)
}
