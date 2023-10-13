package sorter

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Quicksort struct{}

func (qs *Quicksort) Name() string { return "quicksort" }

func (qs *Quicksort) Description() string { return "" }

func (qs *Quicksort) Complexity() Complexity {
	return Complexity{
		TimeBest:   "O(nlogn)",
		TimeAvg:    "O(nlogn)",
		TimeWorst:  "O(n^2)",
		SpaceWorst: "O(logn)",
	}
}

func (qs *Quicksort) WaitForSort(sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg(<-sub)
	}
}

func (qs *Quicksort) Sort(items []Item, sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		sub <- quicksort(items, 0, len(items)-1, sub)
		return FinishMsg{}
	}
}

func quicksort(items []Item, low, high int, sub chan []Item) []Item {
	if low < high {
		var p int
		items, p = partition(items, low, high)
		items = quicksort(items, low, p-1, sub)
		items = quicksort(items, p+1, high, sub)
	}
	return items
}

func partition(items []Item, low, high int) ([]Item, int) {
	pivot := items[high]
	i := low
	for j := low; j < high; j++ {
		if items[j].Value < pivot.Value {
			items[i], items[j] = items[j], items[i]
			i++
		}
	}
	items[i], items[high] = items[high], items[i]
	return items, i
}
