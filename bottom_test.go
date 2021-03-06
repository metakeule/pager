package pager

import (
	"reflect"
	"testing"
)

func TestBottomNext(t *testing.T) {
	tests := []struct {
		times        int
		lines        []string
		selectedLine string
		changed      bool
	}{
		{0, []string{"one", "two", "three"}, "one", false},
		{1, []string{"one", "two", "three"}, "two", true},
		{2, []string{"one", "two", "three"}, "three", true},
		{3, []string{"two", "three", "four"}, "four", true},
		{4, []string{"three", "four", "five"}, "five", true},
		{5, []string{"four", "five", "six"}, "six", true},
		{6, []string{"five", "six", "seven"}, "seven", true},
		{7, []string{"six", "seven", "eight"}, "eight", true},
		{8, []string{"seven", "eight", "nine"}, "nine", true},
		{9, []string{"eight", "nine", "ten"}, "ten", true},
		{10, []string{"eight", "nine", "ten"}, "ten", false},
		{11, []string{"eight", "nine", "ten"}, "ten", false},
	}

	for _, test := range tests {
		pager := newPager(3, Bottom())

		var changed bool

		for i := 0; i < test.times; i++ {
			changed = pager.Next()
		}

		lines, selectedLine := displayData(pager)

		if got, want := selectedLine, test.selectedLine; got != want {
			t.Errorf("%v times p.Next(); selectedLine = %#v; want %#v", test.times, got, want)
		}

		if got, want := lines, test.lines; !reflect.DeepEqual(got, want) {
			t.Errorf("%v times p.Next(); lines = %v; want %v", test.times, got, want)
		}

		if got, want := changed, test.changed; got != want {
			t.Errorf("%v times p.Next(); changed = %v; want %v", test.times, got, want)
		}
	}

}

func TestBottomPrev(t *testing.T) {
	tests := []struct {
		times        int
		lines        []string
		selectedLine string
		changed      bool
	}{
		{11, []string{"one", "two", "three"}, "one", false},
		{10, []string{"one", "two", "three"}, "one", false},
		{9, []string{"one", "two", "three"}, "one", true},
		{8, []string{"one", "two", "three"}, "two", true},
		{7, []string{"one", "two", "three"}, "three", true},
		{6, []string{"two", "three", "four"}, "four", true},
		{5, []string{"three", "four", "five"}, "five", true},
		{4, []string{"four", "five", "six"}, "six", true},
		{3, []string{"five", "six", "seven"}, "seven", true},
		{2, []string{"six", "seven", "eight"}, "eight", true},
		{1, []string{"seven", "eight", "nine"}, "nine", true},
		{0, []string{"eight", "nine", "ten"}, "ten", false},
	}

	for _, test := range tests {
		pager := newPager(3, Bottom())

		var changed bool

		for i := 0; i < 11; i++ {
			pager.Next()
		}

		for i := 0; i < test.times; i++ {
			changed = pager.Prev()
		}

		lines, selectedLine := displayData(pager)

		if got, want := selectedLine, test.selectedLine; got != want {
			t.Errorf("%v times p.Prev(); selectedLine = %#v; want %#v", test.times, got, want)
		}

		if got, want := lines, test.lines; !reflect.DeepEqual(got, want) {
			t.Errorf("%v times p.Prev(); lines = %v; want %v", test.times, got, want)
		}

		if got, want := changed, test.changed; got != want {
			t.Errorf("%v times p.Prev(); changed = %v; want %v", test.times, got, want)
		}
	}

}

func TestBottomPageDown(t *testing.T) {
	tests := []struct {
		times        int
		lines        []string
		selectedLine string
		changed      bool
	}{
		{0, []string{"one", "two", "three"}, "one", false},
		{1, []string{"four", "five", "six"}, "six", true},
		{2, []string{"seven", "eight", "nine"}, "nine", true},
		{3, []string{"eight", "nine", "ten"}, "ten", true},
		{4, []string{"eight", "nine", "ten"}, "ten", false},
	}

	for _, test := range tests {
		pager := newPager(3, Bottom())

		var changes bool

		for i := 0; i < test.times; i++ {
			changes = pager.PageDown()
		}

		lines, selectedLine := displayData(pager)

		if got, want := selectedLine, test.selectedLine; got != want {
			t.Errorf("%v times p.PageDown(); selectedLine = %#v; want %#v", test.times, got, want)
		}

		if got, want := lines, test.lines; !reflect.DeepEqual(got, want) {
			t.Errorf("%v times p.PageDown(); lines = %v; want %v", test.times, got, want)
		}

		if got, want := changes, test.changed; got != want {
			t.Errorf("%v times p.PageDown(); changes = %#v; want %#v", test.times, got, want)
		}
	}

}

func TestBottomPageUp(t *testing.T) {
	tests := []struct {
		timesPageDown int
		timesPageUp   int
		lines         []string
		selectedLine  string
		changed       bool
	}{
		{4, 3, []string{"one", "two", "three"}, "one", true},
		{4, 2, []string{"two", "three", "four"}, "four", true},
		{4, 1, []string{"five", "six", "seven"}, "seven", true},
		{4, 0, []string{"eight", "nine", "ten"}, "ten", false},

		{2, 2, []string{"one", "two", "three"}, "three", true},
		{2, 1, []string{"four", "five", "six"}, "six", true},
		{2, 0, []string{"seven", "eight", "nine"}, "nine", false},
	}

	for _, test := range tests {
		pager := newPager(3, Bottom())

		var changes bool

		for i := 0; i < test.timesPageDown; i++ {
			pager.PageDown()
		}

		for i := 0; i < test.timesPageUp; i++ {
			changes = pager.PageUp()
		}

		lines, selectedLine := displayData(pager)

		if got, want := selectedLine, test.selectedLine; got != want {
			t.Errorf("%v times p.PageDown(); %v times p.PageUp(); selectedLine = %#v; want %#v", test.timesPageDown, test.timesPageUp, got, want)
		}

		if got, want := lines, test.lines; !reflect.DeepEqual(got, want) {
			t.Errorf("%v times p.PageDown(); %v times p.PageUp(); lines = %v; want %v", test.timesPageDown, test.timesPageUp, got, want)
		}

		if got, want := changes, test.changed; got != want {
			t.Errorf("%v times p.PageDown(); %v times p.PageUp(); changes = %#v; want %#v", test.timesPageDown, test.timesPageUp, got, want)
		}
	}

}

func BenchmarkBottom(b *testing.B) {
	b.StopTimer()

	pg := New(40, 50000, Bottom())
	for i := 0; i < 40000; i++ {
		pg.PageUp()
	}

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pg.Indexes()
	}
}
