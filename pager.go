package pager

// type Style func(*pager) (from, to, selected int)
type Option func(*pager)

// Pager allows paging without having to deal
// with the data that is to be paged.
type Pager interface {

	// Next selects the next item. Returns wether the selected item has changed.
	Next() (changed bool)

	// Prev selects the previous item. Returns wether the selected item has changed.
	Prev() (changed bool)

	// PageDown selects the next page. Returns wether the selected item has changed.
	PageDown() (changed bool)

	// PageUp selects the previous page. Returns wether the selected item has changed.
	PageUp() (changed bool)

	// Indexes returns the from, to and selected index.
	// Selected index is counted from start to end, so
	// to get the index of the underlying data, add from to selected.
	// If selected is -1, there is no selection.
	// If from is -1, there is no data.
	Indexes() (from, to, selected int)
}

// PreSelect the given index in the data slice
func PreSelect(index uint) Option {
	return func(pg *pager) {
		pg.selected = int(index)
	}
}

// FixPage sets the fix page display style (default)
func FixPage() Option {
	return func(pg *pager) {

		pg.style = func(p *pager) (from, to, selected int) {

			if p.dataLen == 0 || p.selected > p.dataLen-1 {
				return -1, -1, -1
			}

			page := p.currentPage()
			if page > 0 {
				from = page * p.height
			}
			to = from + p.height
			if p.dataLen < to {
				to = p.dataLen
			}

			return from, to, p.selected - from
		}
	}
}

// Top keeps the selected line at the top
func Top() Option {
	return func(pg *pager) {

		pg.style = func(p *pager) (from, to, selected int) {

			if p.dataLen == 0 || p.selected > p.dataLen-1 {
				return -1, -1, -1
			}
			from = p.selected
			to = p.selected + p.height

			if p.dataLen < to {
				to = p.dataLen
			}

			return from, to, 0
		}
	}
}

// Bottom keeps the selected line at the bottom
func Bottom() Option {
	return func(pg *pager) {

		pg.style = func(p *pager) (from, to, selected int) {

			if p.dataLen == 0 || p.selected > p.dataLen-1 {
				return -1, -1, -1
			}
			if p.selected < p.height {
				to = p.dataLen
				if to > p.height {
					to = p.height
				}
				return 0, to, p.selected
			}

			to = p.selected + 1
			from = to - p.height
			return from, to, p.selected - from
		}
	}
}

type pager struct {
	dataLen, selected, height int
	dataLenDivHeight          int
	style                     func(*pager) (from, to, selected int)
}

// New create a new pager.
// Create a new pager each time the height or dataLen changes.
//func New(height, dataLen, selected int, style Style) Pager {
func New(height, dataLen int, opts ...Option) Pager {
	p := &pager{height: height, dataLen: dataLen}
	p.dataLenDivHeight = dataLen / height

	for _, opt := range opts {
		opt(p)
	}

	if p.style == nil {
		FixPage()(p)
	}

	if dataLen == 0 {
		p.selected = -1
	}
	return p
}

// Next selects the next item. Returns wether the selected item has changed.
func (p *pager) Next() (changed bool) {
	if p.dataLen-1 > p.selected {
		p.selected++
		changed = true
	}
	return
}

// Prev selects the previous item. Returns wether the selected item has changed.
func (p *pager) Prev() (changed bool) {
	if p.selected > 0 {
		p.selected--
		changed = true
	}
	return
}

// PageDown selects the next page. Returns wether the selected item has changed.
func (p *pager) PageDown() (changed bool) {
	if p.dataLen == 0 {
		return
	}

	page := p.currentPage()
	if page >= p.dataLenDivHeight {
		return
	}

	page++
	p.selected = (page+1)*p.height - 1
	changed = true

	if p.selected > p.dataLen-1 {
		p.selected = p.dataLen - 1
	}
	return
}

// PageUp selects the previous page. Returns wether the selected item has changed.
func (p *pager) PageUp() (changed bool) {
	if p.currentPage() == 0 {
		return
	}

	p.selected -= p.height
	changed = true
	return
}

// Indexes returns the from, to and selected index.
// Selected index is counted from start to end, so
// to get the index of the underlying data, add from to selected.
// If selected is -1, there is no selection.
// If from is -1, there is no data.
func (p *pager) Indexes() (from, to, selected int) {
	return p.style(p)
}

func (p *pager) currentPage() (page int) {
	if p.selected < 0 {
		return
	}

	page = p.selected / p.height
	return
}
