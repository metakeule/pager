
pager
=====

[![Build Status](https://secure.travis-ci.org/metakeule/pager.png)](http://travis-ci.org/metakeule/pager)

100% test coverage (that was easy :-))

`pager` provides a simple data neutral paging solution for Go


Usage
-----

```go
package main

import (
	"fmt"
	"github.com/metakeule/pager"
)

var data = []string{"one", "two", "three", "four", "five", "six"}

func print(pg pager.Pager) {
	from, to, selected := pg.Indexes()

    // we got data
	if from > -1 {  
		for i, line := range data[from:to] {
			prefix := "  "
			if i == selected {
				prefix = "> "
			}
			fmt.Println(prefix + line)
		}
	}
}

func main() {
	
	pg := pager.New(3, len(data))

    // do the paging stuff here
    // when height or data changes, create a new pager
	pg.PageDown()
	pg.Prev()

    // print the paged data
	print(pg)

	//   three
	//   four
	// > five
}
```

Documentation
-------------

see http://godoc.org/github.com/metakeule/pager