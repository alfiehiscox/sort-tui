package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	Value int
	Bar   string
	Color lipgloss.Color
}

type SwapMsg struct {
	IndexA int
	IndexB int
}

type FinishMsg bool

type Sorter interface {
	Sort([]int, func(tea.Msg))
	Code() string
	Description() string
	References() []string
}

type InsertionSort struct{}

func (is *InsertionSort) Sort(items []Item, update func(tea.Msg)) {
	for i := range items {
		for j := i; j > 0 && items[j-1].Value > items[j].Value; j-- {
			time.Sleep(100 * time.Millisecond)
			items[j], items[j-1] = items[j-1], items[j]
			update(SwapMsg{IndexA: j, IndexB: j - 1})
		}
	}
	update(FinishMsg(true))
}

func (is *InsertionSort) Code() string         { return "INSERTION CODE" }
func (is *InsertionSort) Description() string  { return "INSERTION DISCRIPTION" }
func (is *InsertionSort) References() []string { return []string{"INSERTION REF"} }
