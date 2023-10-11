package sorter

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	ICON          = '■'
	TIME_INTERVAL = 100 * time.Millisecond
)

var Sorters = []Sorter{
	&InsertionSort{},
	&BubbleSort{},
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

type FinishMsg struct{} // Passed on finish
type UpdateMsg []Item   // Passed on each move/swap

type Sorter interface {
	Name() string
	Description() string
	Complexity() Complexity

	Sort([]Item, chan []Item) tea.Cmd
	WaitForSort(chan []Item) tea.Cmd
}

// ========== UTILS ============ //

// Add numbers to a string, until the len of the string is
// greater than the maxCell

// Returns the number of elements possible to fit into a
// contingous group of cells, separated by commas with no
// trailing commas.
// Example: MaxArrayFromCells(10) -> 1,2,3,4,5 -> 9 cells -> 5 elements
// Example: MaxArrayFromCells(35) == 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15 == 35 cells == 15 elements
func MaxArrayFromCells(maxCell int) int {
	var buf bytes.Buffer
	comma := true // Was the last added byte a comma
	count := 1
	for i := 1; buf.Len() < maxCell; i++ {
		if comma {
			buf.WriteString(fmt.Sprintf("%d", count)) // Add number
			count++
			comma = false
		} else {
			buf.WriteByte(',') // Add comma
			comma = true
		}
	}
	s := buf.String()
	if comma { // trailing comma
		s = s[:len(s)-1]
	}
	return len(strings.Split(s, ","))
}

func GetRandomItems(size int) []Item {
	var items []Item
	values := rand.Perm(size)
	for _, value := range values {
		items = append(items, Item{
			Value:   value + 1,
			Bar:     makeBar(value + 1),
			Focused: false,
		})
	}
	return items
}

func makeBar(size int) string {
	var buf bytes.Buffer
	s := fmt.Sprintf("%d", size)
	buf.WriteString(s)
	for i := 0; i < 3-len(s); i++ {
		buf.WriteByte(' ')
	}
	for i := 0; i < size*2; i++ {
		buf.WriteRune(ICON)
	}
	return buf.String()
}
