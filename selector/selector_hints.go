package selector

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var hints = widgets.NewList()

func init() {
	hints = widgets.NewList()
	hints.Title = "Hints"
	hints.TitleStyle.Fg = termui.ColorCyan
	hints.Rows = hintsStr
	hints.TextStyle.Fg = termui.ColorYellow
	hints.Border = true
	hints.BorderStyle.Fg = termui.ColorCyan
}

func buildHints() termui.Drawable {
	hints.SetRect(0, 3, sidebarWidth, termHeight())

	return hints
}
