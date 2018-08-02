package pager_test

import (
	"fmt"
	"github.com/metakeule/pager"
)

func Example() {
	fmt.Println("")

	var data = []string{"one", "two", "three", "four", "five", "six", "seven"}

	pg := pager.New(3, len(data), pager.PreSelect(3), pager.Top())

	pg.PageDown()
	pg.Prev()
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
	// > five
	//   six
	//   seven
}
