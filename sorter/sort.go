package sorter

import (
	"bytes"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const ICON = '#'

type SorterType string

var Sorters = []Sorter{
	&InsertionSort{},
	&AlfieSort{},
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

type UpdateMsg []Item
type FinishMsg bool

type Sorter interface {
	Name() string
	Description() string
	Sort([]Item, func(tea.Msg))
	Code() string
	Complexity() Complexity
}

type InsertionSort struct{}

func (is *InsertionSort) Sort(items []Item, update func(tea.Msg)) {
	for i := range items {
		for j := i; j > 0 && items[j-1].Value > items[j].Value; j-- {
			items[j], items[j-1] = items[j-1], items[j]
			items[j-1].Focused = true
			update(UpdateMsg(items))
			time.Sleep(100 * time.Millisecond)
			items[j-1].Focused = false
		}
	}
	update(FinishMsg(true))
}

func (is *InsertionSort) Description() string { return "INSERTION DISCRIPTION" }
func (is *InsertionSort) Name() string        { return "insertionsort" }
func (is *InsertionSort) Code() string        { return "INSERTION CODE" }
func (is *InsertionSort) Complexity() Complexity {
	return Complexity{
		TimeBest:   "O(n)",
		TimeWorst:  "O(n^2)",
		TimeAvg:    "O(n^2)",
		SpaceWorst: "O(1)",
	}
}

type AlfieSort struct{}

func (is *AlfieSort) Sort(items []Item, update func(tea.Msg)) {
	for i := range items {
		for j := i; j > 0 && items[j-1].Value > items[j].Value; j-- {
			items[j], items[j-1] = items[j-1], items[j]
			items[j-1].Focused = true
			update(UpdateMsg(items))
			time.Sleep(100 * time.Millisecond)
			items[j-1].Focused = false
		}
	}
	update(FinishMsg(true))
}

func (is *AlfieSort) Description() string { return "INSERTION DISCRIPTION" }
func (is *AlfieSort) Name() string        { return "alfiesort" }
func (is *AlfieSort) Code() string        { return "INSERTION CODE" }
func (is *AlfieSort) Complexity() Complexity {
	return Complexity{
		TimeBest:   "O(n^6)",
		TimeWorst:  "O(n^n)",
		TimeAvg:    "O(e)",
		SpaceWorst: "O(!)",
	}
}

func makeBar(size int) string {
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteRune(ICON)
	}
	return buf.String()
}
