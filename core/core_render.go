package core

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Render refresh all ui
func Render() {
	ui.Clear()
	ui.Render(drawServerTable())
}

// draw server table
func drawServerTable() ui.Drawable {
	table3 := widgets.NewTable()
	table3.Rows = [][]string{
		{"header1", "header2", "header3"},
		{"AAA", "BBB", "CCC"},
		{"DDD", "EEE", "FFF"},
		{"GGG", "HHH", "III"},
	}
	table3.TextStyle = ui.NewStyle(ui.ColorWhite)
	table3.RowSeparator = true
	table3.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table3.SetRect(0, 30, 70, 20)
	table3.FillRow = true
	table3.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table3.RowStyles[2] = ui.NewStyle(ui.ColorWhite, ui.ColorRed, ui.ModifierBold)
	table3.RowStyles[3] = ui.NewStyle(ui.ColorYellow)
	return table3
}
