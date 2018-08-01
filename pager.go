package pager

// Pager allows paging without having to deal
// with the data that is to be paged.
type Pager interface {
	Next() (changed bool)
	Prev() (changed bool)
	PageDown() (changed bool)
	PageUp() (changed bool)
	Indexes() (from, to, selected int)
}

type pager struct {
	dataLen, selected, height int
	dataLenDivHeight          int
}

// New create a new pager.
// Create a new pager each time the height or dataLen changes.
func New(height, dataLen int) Pager {
	p := &pager{height: height, dataLen: dataLen}
	if dataLen == 0 {
		p.selected = -1
	}
	p.dataLenDivHeight = dataLen / height
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

func (p *pager) currentPage() (page int) {
	if p.selected < 0 {
		return
	}

	page = p.selected / p.height
	return
}
