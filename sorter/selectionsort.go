package sorter

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectionSort struct{}

func (ss *SelectionSort) Name() string { return "selectionsort" }

func (ss *SelectionSort) Description() string {
	return "Selection sort is an in place comparison sorting algorithm. It has quadratic time complexity " +
		"which makes it inefficient on large lists, and generally performs worse than the similar insertion sort. " +
		"Selection sort is noted for its simplicity and has performance advantages over more complicated algorithms " +
		"in certain situations, particularly where auxiliary memory is limited.\n\n" +
		"The algorithm divides the input list into two parts: a sorted sublist of items which is built up from left to " +
		"right at the front of the list and a sublist of the remaining unsorted items that occupy the rest of the list. " +
		"Initially, the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm proceeds " +
		"by finding the smallest element in the unsorted sublist, swapping it with the leftmost unsorted element " +
		"and moving the sublist boundaries one element to the right."
}

func (ss *SelectionSort) Complexity() Complexity {
	return Complexity{
		TimeBest:   "O(n^2)",
		TimeAvg:    "O(n^2)",
		TimeWorst:  "O(n^2)",
		SpaceWorst: "O(1)",
	}
}

func (ss *SelectionSort) WaitForSort(sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg(<-sub)
	}
}

func (ss *SelectionSort) Sort(items []Item, sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		for i := 0; i < len(items)-1; i++ {
			items[i].Focused = true
			min := i
			for j := i + 1; j < len(items); j++ {
				items[j].Focused = true
				if items[j].Value < items[min].Value {
					min = j
				}
				sub <- items
				time.Sleep(TIME_INTERVAL)
				items[j].Focused = false
			}
			if min != i {
				items[i], items[min] = items[min], items[i]
			}
			items[i].Focused = false
			items[min].Focused = false
		}
		return FinishMsg{}
	}
}
