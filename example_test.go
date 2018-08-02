package pager_test

import (
	"fmt"
	"github.com/metakeule/pager"
)

func Example() {
	fmt.Println("")

	var data = []string{"one", "two", "three", "four", "five", "six"}

	height := 3
	selected := 0
	pg := pager.New(height, len(data), selected)

	pg.PageDown()
	pg.Prev()

	from, to, selected := pg.Indexes()

	if from > -1 {
		for i, line := range data[from:to] {
			prefix := "  "
			if i == selected {
				prefix = "> "
			}
			fmt.Println(prefix + line)
		}
	}

	// Output:
	//   three
	//   four
	// > five
}
