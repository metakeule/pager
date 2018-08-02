package pager

import (
	"reflect"
	"testing"
)

func TestTopNext(t *testing.T) {
	tests := []struct {
		times        int
		lines        []string
		selectedLine string
		changed      bool
	}{
		{0, []string{"one", "two", "three"}, "one", false},
		{1, []string{"two", "three", "four"}, "two", true},
		{2, []string{"three", "four", "five"}, "three", true},
		{3, []string{"four", "five", "six"}, "four", true},
		{4, []string{"five", "six", "seven"}, "five", true},
		{5, []string{"six", "seven", "eight"}, "six", true},
		{6, []string{"seven", "eight", "nine"}, "seven", true},
		{7, []string{"eight", "nine", "ten"}, "eight", true},
		{8, []string{"nine", "ten"}, "nine", true},
		{9, []string{"ten"}, "ten", true},
		{10, []string{"ten"}, "ten", false},
		{11, []string{"ten"}, "ten", false},
	}

	for _, test := range tests {
		pager := newPager(3, Top())

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

func TestTopPrev(t *testing.T) {
	tests := []struct {
		times        int
		lines        []string
		selectedLine string
		changed      bool
	}{
		{11, []string{"one", "two", "three"}, "one", false},
		{10, []string{"one", "two", "three"}, "one", false},
		{9, []string{"one", "two", "three"}, "one", true},
		{8, []string{"two", "three", "four"}, "two", true},
		{7, []string{"three", "four", "five"}, "three", true},
		{6, []string{"four", "five", "six"}, "four", true},
		{5, []string{"five", "six", "seven"}, "five", true},
		{4, []string{"six", "seven", "eight"}, "six", true},
		{3, []string{"seven", "eight", "nine"}, "seven", true},
		{2, []string{"eight", "nine", "ten"}, "eight", true},
		{1, []string{"nine", "ten"}, "nine", true},
		{0, []string{"ten"}, "ten", false},
	}

	for _, test := range tests {
		pager := newPager(3, Top())

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

func TestTopPageDown(t *testing.T) {
	tests := []struct {
		times        int
		lines        []string
		selectedLine string
		changed      bool
	}{
		{0, []string{"one", "two", "three"}, "one", false},
		{1, []string{"six", "seven", "eight"}, "six", true},
		{2, []string{"nine", "ten"}, "nine", true},
		{3, []string{"ten"}, "ten", true},
		{4, []string{"ten"}, "ten", false},
	}

	for _, test := range tests {
		pager := newPager(3, Top())

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

func TestTopPageUp(t *testing.T) {
	tests := []struct {
		timesPageDown int
		timesPageUp   int
		lines         []string
		selectedLine  string
		changed       bool
	}{
		{4, 3, []string{"one", "two", "three"}, "one", true},
		{4, 2, []string{"four", "five", "six"}, "four", true},
		{4, 1, []string{"seven", "eight", "nine"}, "seven", true},
		{4, 0, []string{"ten"}, "ten", false},

		{2, 2, []string{"three", "four", "five"}, "three", true},
		{2, 1, []string{"six", "seven", "eight"}, "six", true},
		{2, 0, []string{"nine", "ten"}, "nine", false},
	}

	for _, test := range tests {
		pager := newPager(3, Top())

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

func BenchmarkTop(b *testing.B) {
	b.StopTimer()

	pg := New(40, 50000, Top())
	for i := 0; i < 40000; i++ {
		pg.PageUp()
	}

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pg.Indexes()
	}
}
