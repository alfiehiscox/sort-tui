package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const SIZE = 20
const ICON = '#'

func main() {
	ix := getRandomItems()
	ixc := make([]Item, len(ix))
	copy(ixc, ix)

	m := Model{items: ix}
	p := tea.NewProgram(m)
	sort := InsertionSort{}
	go sort.Sort(ixc, func(msg tea.Msg) {
		p.Send(msg)
	})

	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}

func getRandomItems() []Item {
	var items []Item
	values := rand.Perm(SIZE)
	for _, value := range values {
		items = append(items, Item{
			Value: value,
			Bar:   makeBar(value),
			Color: lipgloss.Color("1"),
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
