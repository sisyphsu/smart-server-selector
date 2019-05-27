package selector

import (
	"github.com/gizak/termui"
)

// Draw refresh all ui
func Draw() {
	//termui.Clear()
	termui.Render(drawServerTable())
}

// draw server table
func drawServerTable() (table *termui.Table) {
	data := cxt.buildTableData()
	table = termui.NewTable()
	table.Rows = data
	table.FgColor = termui.ColorGreen
	table.BgColor = termui.ColorDefault
	table.TextAlign = termui.AlignLeft
	table.Separator = false
	table.Analysis()
	table.CellWidth = []int{4, 8, termui.TermWidth() - sidebarWidth - 12}
	table.SetSize()
	table.Y = 3
	table.X = sidebarWidth
	table.Border = true
	table.BorderFg = termui.ColorCyan
	table.BorderLabel = "Servers"
	table.BorderLabelFg = termui.ColorCyan
	table.Width = termui.TermWidth() - sidebarWidth
	// handle selected
	for i, row := range data {
		if row[1] != cxt.serverSelected {
			continue
		}
		table.BgColors[i] = termui.ColorRed
		table.FgColors[i] = termui.ColorYellow
	}
	return
}
