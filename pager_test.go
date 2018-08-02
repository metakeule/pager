package pager

import (
	"testing"
)

func newPager(height int, opts ...Option) *pager {
	p := New(height, len(data), opts...)
	return p.(*pager)
}

var data = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func displayData(pg Pager) (lines []string, selectedLine string) {
	from, to, selected := pg.Indexes()

	if from == -1 {
		return
	}

	if selected != -1 {
		selectedLine = data[selected+from]
	}

	lines = data[from:to]
	return
}

func TestCurrentPage(t *testing.T) {
	tests := []struct {
		height   int
		selected int
		page     int
	}{
		{3, 0, 0},
		{3, 1, 0},
		{3, 2, 0},
		{3, 3, 1},
		{3, 4, 1},
		{3, 5, 1},
		{3, 6, 2},
		{3, 7, 2},
		{3, 8, 2},
		{3, 9, 3},
	}

	for _, test := range tests {
		pager := newPager(test.height)
		pager.selected = test.selected

		page := pager.currentPage()

		if got, want := page, test.page; got != want {
			t.Errorf("height: %v selected: %v; page = %#v; want %#v", test.height, test.selected, got, want)
		}

	}

}

func TestEmpty(t *testing.T) {
	pg := New(3, 0)

	from, to, selected := pg.Indexes()

	if from != -1 || to != -1 || selected != -1 {
		t.Errorf("from: %v, to: %v, selected: %v", from, to, selected)
	}

	pg.PageDown()

	from, to, selected = pg.Indexes()

	if from != -1 || to != -1 || selected != -1 {
		t.Errorf("after PageDown: from: %v, to: %v, selected: %v", from, to, selected)
	}

	pg.PageUp()

	from, to, selected = pg.Indexes()

	if from != -1 || to != -1 || selected != -1 {
		t.Errorf("after PageUp: from: %v, to: %v, selected: %v", from, to, selected)
	}
}

func BenchmarkNext(b *testing.B) {
	b.StopTimer()

	pg := New(40, 50000)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pg.Next()
	}
}

func BenchmarkPrev(b *testing.B) {
	b.StopTimer()

	pg := New(40, 50000)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pg.Prev()
	}
}

func BenchmarkPageDown(b *testing.B) {
	b.StopTimer()

	pg := New(40, 50000)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pg.PageDown()
	}
}

func BenchmarkPageUp(b *testing.B) {
	b.StopTimer()

	pg := New(40, 50000)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pg.PageUp()
	}
}
