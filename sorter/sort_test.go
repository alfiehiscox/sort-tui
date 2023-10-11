package sorter

import "testing"

func TestMaxArrayFromCells(t *testing.T) {
	tests := []struct {
		Input  int
		Output int
	}{
		{10, 5},
		{35, 15},
		{50, 20},
		{100, 37},
	}

	for _, test := range tests {
		r := MaxArrayFromCells(test.Input)
		if r != test.Output {
			t.Fatalf("Expected=%d,Got=%d", test.Output, r)
		}
	}
}
