package sorter

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type BubbleSort struct{}

func (bs *BubbleSort) Name() string { return "bubblesort" }

func (bs *BubbleSort) Description() string {
	return "Bubble sort, sometimes referred to as sinking sort, is a simple sorting " +
		"algorithm that repeatedly steps through the input list element by element," +
		"comparing the current element with the one after it, swapping their values if " +
		"needed. These passes through the list are repeated until no swaps had to be " +
		"performed during a pass, meaning that the list has become fully sorted.\n\n" +
		"This simple algorithm performs poorly in real world use and is used primarily " +
		"as an educational tool. More efficient algorithms such as quicksort, timsort, or " +
		"merge sort are used by the sorting libraries built into popular programming languages " +
		"such as Python and Java. However, if parallel processing is allowed, bubble sort " +
		"sorts in O(n) time, making it considerably faster than parallel implementations of " +
		"insertion sort or selection sort which do not parallelize as effectively."
}

func (bs *BubbleSort) Complexity() Complexity {
	return Complexity{
		TimeBest:   "O(n)",
		TimeAvg:    "O(n^2)",
		TimeWorst:  "O(n^2)",
		SpaceWorst: "O(1)",
	}
}

func (bs *BubbleSort) Sort(items []Item, sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		swapped := false
		for i := 0; i < len(items)-1; i++ {
			swapped = false
			for j := 0; j < len(items)-1; j++ {
				if items[j].Value > items[j+1].Value {
					items[j], items[j+1] = items[j+1], items[j]
					items[j+1].Focused = true
					sub <- items
					swapped = true
					time.Sleep(TIME_INTERVAL)
					items[j+1].Focused = false
				}
			}
			if !swapped {
				break
			}
		}
		return FinishMsg{}
	}
}

func (bs *BubbleSort) WaitForSort(sub chan []Item) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg(<-sub)
	}
}
