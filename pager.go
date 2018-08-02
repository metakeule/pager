package pager

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

type pager struct {
	dataLen, selected, height int
	dataLenDivHeight          int
	style                     func(*pager) (from, to, selected int)
}

// New create a new pager.
// Each time the height or dataLen changes, a new pager should be created.
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
	if p.dataLen == 0 || p.selected > p.dataLen-1 {
		return -1, -1, -1
	}

	return p.style(p)
}

func (p *pager) currentPage() (page int) {
	if p.selected < 0 {
		return
	}

	page = p.selected / p.height
	return
}
