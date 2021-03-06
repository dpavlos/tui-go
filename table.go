package tui

import "image"

var _ Widget = &Table{}

// Table is a widget that lays out widgets in a table.
type Table struct {
	selected           int
	onItemActivated    func(*Table)
	onSelectionChanged func(*Table)
	headers            []string

	*Grid
}

// Table returns a new Table.
func NewTable(cols, rows int) *Table {
	return &Table{
		Grid: NewGrid(cols, rows),
	}
}

// Draw draws the table.
func (g *Table) Draw(p *Painter) {
	s := g.Size()

	if g.hasBorder {
		border := 1

		// Draw outmost border.
		p.DrawRect(0, 0, s.X, s.Y)

		// Draw column dividers.
		var coloff int
		for i := 0; i < g.cols-1; i++ {
			x := g.colWidths[i] + coloff + border
			p.DrawVerticalLine(x, 0, s.Y-1)
			p.DrawRune(x, 0, '┬')
			p.DrawRune(x, s.Y-1, '┴')
			coloff = x
		}

		// Draw row dividers.
		var rowoff int
		for j := 0; j < g.rows-1; j++ {
			y := g.rowHeights[j] + rowoff + border
			p.DrawHorizontalLine(0, s.X-1, y)
			p.DrawRune(0, y, '├')
			p.DrawRune(s.X-1, y, '┤')
			rowoff = y
		}

		// Polish the intersections.
		rowoff = 0
		for j := 0; j < g.rows-1; j++ {
			y := g.rowHeights[j] + rowoff + border
			coloff = 0
			for i := 0; i < g.cols-1; i++ {
				x := g.colWidths[i] + coloff + border
				p.DrawRune(x, y, '┼')
				coloff = x
			}
			rowoff = y
		}
	}

	// Draw cell content.
	for i := 0; i < g.cols; i++ {
		for j := 0; j < g.rows; j++ {
			style := "table.cell"
			if j == g.selected {
				style += ".selected"
			}

			p.WithStyledBrush(style, func(p *Painter) {
				pos := image.Point{i, j}
				wp := g.mapCellToLocal(pos)
				w := g.colWidths[i]

				p.FillRect(wp.X, wp.Y, w, 1)

				if w, ok := g.cells[pos]; ok {
					p.Translate(wp.X, wp.Y)
					w.Draw(p)
					p.Restore()
				}
			})
		}
	}
}

func (t *Table) OnEvent(ev Event) {
	switch ev.Key {
	case KeyArrowUp:
		t.moveUp()
	case KeyArrowDown:
		t.moveDown()
	case KeyEnter:
		if t.onItemActivated != nil {
			t.onItemActivated(t)
		}
	}

	switch ev.Ch {
	case 'k':
		t.moveUp()
	case 'j':
		t.moveDown()
	}
}

func (t *Table) moveUp() {
	if t.selected > 0 {
		t.selected--
	}
	if t.onSelectionChanged != nil {
		t.onSelectionChanged(t)
	}
}

func (t *Table) moveDown() {
	if t.selected < t.rows-1 {
		t.selected++
	}
	if t.onSelectionChanged != nil {
		t.onSelectionChanged(t)
	}
}

// SetSelected changes the currently selected item.
func (t *Table) SetSelected(i int) {
	t.selected = i
}

// Selected returns the index of the currently selected item.
func (t *Table) Selected() int {
	return t.selected
}

// Select calls SetSelected and the OnSelectionChanged function.
func (t *Table) Select(i int) {
	t.SetSelected(i)
	t.onSelectionChanged(t)
}

func (t *Table) OnItemActivated(fn func(*Table)) {
	t.onItemActivated = fn
}

func (t *Table) OnSelectionChanged(fn func(*Table)) {
	t.onSelectionChanged = fn
}

func (t *Table) SetHeaders(headers ...string) {
	t.headers = headers
}
