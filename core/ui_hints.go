package core

import "github.com/gizak/termui"

// draw hints
func NewHints() {
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

	hints := termui.NewList()
	hints.Items = strs
	hints.ItemFgColor = termui.ColorYellow
	hints.BorderLabel = "Hints"
	hints.BorderLabelFg = termui.ColorCyan
	hints.BorderFg = termui.ColorCyan
	hints.Y = 3
	hints.X = 0
	hints.Width = sidebarWidth
	hints.Height = termui.TermHeight() - 3

	termui.Render(hints)
}
