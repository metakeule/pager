package pager_test

import (
	"fmt"
	"github.com/metakeule/pager"
)

func Example() {
	var data = []string{"one", "two", "three", "four", "five", "six"}

	pg := pager.New(3, len(data))

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

	// Output:   three
	//   four
	// > five
}
