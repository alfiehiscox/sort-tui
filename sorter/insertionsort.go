package sorter

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type InsertionSort struct{}

func (is *InsertionSort) Sort(items []Item, sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		for i := range items {
			for j := i; j > 0 && items[j-1].Value > items[j].Value; j-- {
				items[j], items[j-1] = items[j-1], items[j]
				items[j-1].Focused = true
				sub <- items
				time.Sleep(TIME_INTERVAL)
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
	s += "1) Simple Implementation \n"
	s += "2) Efficient over small data sets \n"
	s += "3) Stable: preserves order of duplicate keys."
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
