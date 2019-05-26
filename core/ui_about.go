package core

import "github.com/gizak/termui"

// draw the logo
func NewAbout() {
	p := termui.NewPar("Smart Server Selector")
	p.X = 0
	p.Height = 3
	p.Width = sidebarWidth
	p.TextFgColor = termui.ColorYellow
	p.BorderLabel = "About"
	p.BorderLabelFg = termui.ColorCyan
	p.BorderFg = termui.ColorCyan

	termui.Render(p)
}
