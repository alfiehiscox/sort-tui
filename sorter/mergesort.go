package sorter

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type MergeSort struct {
	copy []Item
}

func (ms *MergeSort) Name() string           { return "mergesort" }
func (ms *MergeSort) Description() string    { return "" }
func (ms *MergeSort) Complexity() Complexity { return Complexity{} }

func (ms *MergeSort) WaitForSort(sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg(<-sub)
	}
}

func (ms *MergeSort) Sort(items []Item, sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		c := make([]Item, len(items))
		copy(c, items)
		ms.copy = c
		ms.mergesort(items, sub)
		return FinishMsg{}
	}
}

func (ms *MergeSort) mergesort(items []Item, sub chan []Item) []Item {
	if len(items) <= 1 {
		return items
	}
	mid := len(items) / 2
	left := ms.mergesort(items[:mid], sub)
	right := ms.mergesort(items[mid:], sub)
	return ms.merge(left, right, sub)
}

func (ms *MergeSort) merge(left, right []Item, sub chan []Item) []Item {
	test := make([]Item, len(left)+len(right))
	test = append(test, left...)
	test = append(test, right...)
	offset := ms.findFirst(test)

	result := make([]Item, len(left)+len(right))
	i, j := 0, 0

	for k := 0; k < len(result); k++ {
		if i >= len(left) {
			result[k] = right[j]
			ms.hotswappy(right[j], offset+k)
			sub <- ms.copy
			j++
		} else if j >= len(right) {
			result[k] = left[i]
			ms.hotswappy(left[i], offset+k)
			sub <- ms.copy
			i++
		} else if left[i].Value < right[j].Value {
			result[k] = left[i]
			ms.hotswappy(left[i], offset+k)
			sub <- ms.copy
			i++
		} else {
			result[k] = right[j]
			ms.hotswappy(right[j], offset+k)
			sub <- ms.copy
			j++
		}
		time.Sleep(TIME_INTERVAL / 2)
	}

	return result
}

func (ms *MergeSort) findFirst(results []Item) int {
	for index, item := range ms.copy {
		for _, r := range results {
			if item == r {
				return index
			}
		}
	}
	return -1
}

func (ms *MergeSort) hotswappy(result Item, target int) {
	for index, item := range ms.copy {
		if item == result {
			ms.copy[index], ms.copy[target] = ms.copy[target], ms.copy[index]
		}
	}
}
