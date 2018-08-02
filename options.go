package pager

// Option is an option for the page initialization
type Option func(*pager)

// PreSelect the given index in the data slice
func PreSelect(index uint) Option {
	return func(pg *pager) {
		pg.selected = int(index)
	}
}

// FixPage always keeps the same pages (default)
func FixPage() Option {
	return func(pg *pager) {

		pg.style = func(p *pager) (from, to, selected int) {

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
