
pager
=====

[![Build Status Travis](https://secure.travis-ci.org/metakeule/pager.png)](http://travis-ci.org/metakeule/pager) [![Build status Appveyor](https://ci.appveyor.com/api/projects/status/82i4pu5giscl7b13?svg=true)](https://ci.appveyor.com/project/metakeule/pager) [![Go Report](https://goreportcard.com/badge/github.com/metakeule/pager)](https://goreportcard.com/report/github.com/metakeule/pager) [![Coverage Status](https://coveralls.io/repos/github/metakeule/pager/badge.svg?branch=master)](https://coveralls.io/github/metakeule/pager?branch=master)

`pager` provides a simple data neutral paging solution for Go.

- This package does not depend on external packages.
- All versions of Go are supported.

Status
------

Stable

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
	height := 3
	selected := 0
	pg := pager.New(height, len(data), selected)

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

Benchmarks
----------

```
BenchmarkFixPage-4   26.7  ns/op   0 B/op     0 allocs/op
BenchmarkTop-4        9.08 ns/op   0 B/op     0 allocs/op
BenchmarkBottom-4     8.17 ns/op   0 B/op     0 allocs/op

BenchmarkNext-4       4.04 ns/op   0 B/op     0 allocs/op
BenchmarkPrev-4       4.54 ns/op   0 B/op     0 allocs/op
BenchmarkPageDown-4   20.6 ns/op   0 B/op     0 allocs/op
BenchmarkPageUp-4     21.3 ns/op   0 B/op     0 allocs/op
```