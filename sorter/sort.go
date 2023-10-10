package sorter

import (
	"bytes"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const ICON = '#'

type SorterType string

var Sorters = []Sorter{
	&InsertionSort{},
	// &AlfieSort{},
}

type Item struct {
	Value   int
	Bar     string
	Focused bool
}

type Complexity struct {
	TimeBest   string
	TimeAvg    string
	TimeWorst  string
	SpaceWorst string
}

type FinishMsg struct{}
type UpdateMsg []Item

type Sorter interface {
	Name() string
	Description() string
	Sort([]Item, chan []Item) tea.Cmd
	WaitForSort(chan []Item) tea.Cmd
	Complexity() Complexity
}

type InsertionSort struct{}

func (is *InsertionSort) Sort(items []Item, sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		for i := range items {
			for j := i; j > 0 && items[j-1].Value > items[j].Value; j-- {
				items[j], items[j-1] = items[j-1], items[j]
				items[j-1].Focused = true
				sub <- items
				time.Sleep(100 * time.Millisecond)
				items[j-1].Focused = false
			}
		}
		return FinishMsg{}
	}
}

func (is *InsertionSort) WaitForSort(sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg(<-sub)
	}
}

func (is *InsertionSort) Description() string {
	s := "Insertion sort is a simple sorting algorithm that builds the final sorted array "
	s += "(or list) one item at a time by comparisons. It is much less efficient on large "
	s += "lists than more advanced algorithms such as quicksort, heapsort, or merge sort. "
	s += "However, insertion sort provides several advantages: \n\n"
	s += "1) Simple Implementation, \n"
	s += "2) Efficient over small data sets, \n"
	s += "3) Stable: preserve same key order."
	return s
}

func (is *InsertionSort) Name() string { return "insertionsort" }
func (is *InsertionSort) Complexity() Complexity {
	return Complexity{
		TimeBest:   "O(n)",
		TimeWorst:  "O(n^2)",
		TimeAvg:    "O(n^2)",
		SpaceWorst: "O(1)",
	}
}

func GetRandomItems(size int) []Item {
	var items []Item
	values := rand.Perm(size)
	for i, value := range values {
		items = append(items, Item{
			Value:   value + 1,
			Bar:     makeBar(value),
			Focused: i == 0,
		})
	}
	return items
}

func makeBar(size int) string {
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteRune(ICON)
	}
	return buf.String()
}
