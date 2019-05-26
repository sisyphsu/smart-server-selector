package core

import (
	"github.com/gizak/termui"
)

const sidebarWidth = 23

// Draw refresh all ui
func Draw() {
	//termui.Clear()
	termui.Render(drawLogo(), drawHints(), drawServerTable())
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

// draw the logo
func drawLogo() (p *termui.Par) {
	p = termui.NewPar("Smart Server Selector")
	p.X = 0
	p.Height = 3
	p.Width = sidebarWidth
	p.TextFgColor = termui.ColorYellow
	p.BorderLabel = "About"
	p.BorderLabelFg = termui.ColorCyan
	p.BorderFg = termui.ColorCyan
	return
}

//// draw the searchBar
//func drawSearchbar() (p *termui.Par) {
//	label := cxt.keyword
//	if time.Now().Second()%2 == 0 {
//		label += "▂"
//	}
//	p = termui.NewPar(label)
//	p.X = sidebarWidth
//	p.Height = 3
//	p.Width = termui.TermWidth() - sidebarWidth
//	p.TextFgColor = termui.ColorGreen
//	p.BorderLabel = "Search"
//	p.BorderLabelFg = termui.ColorCyan
//	p.BorderFg = termui.ColorCyan
//	return
//}

// draw hints
func drawHints() (ls *termui.List) {
	strs := []string{
		"  [Ctrl+C] exit",
		"  [Ctrl+P] edit",
		"      [Up] prev",
		"    [Down] next",
		"     [Tab] switch",
		"     [Esc] clear",
		"     [Del] delete",
		"   [Enter] confirm",
		"      [....](fg-white)    ",
	}

	ls = termui.NewList()
	ls.Items = strs
	ls.ItemFgColor = termui.ColorYellow
	ls.BorderLabel = "Hints"
	ls.BorderLabelFg = termui.ColorCyan
	ls.BorderFg = termui.ColorCyan
	ls.Y = 3
	ls.X = 0
	ls.Width = sidebarWidth
	ls.Height = termui.TermHeight() - 3
	return ls
}
