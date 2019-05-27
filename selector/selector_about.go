package selector

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var about = widgets.NewParagraph()

func init() {
	about = widgets.NewParagraph()
	about.Title = "About"
	about.TitleStyle.Fg = termui.ColorCyan
	about.Text = "Smart Server Selector"
	about.TextStyle.Fg = termui.ColorYellow
	about.Border = true
	about.BorderStyle.Fg = termui.ColorCyan
}

func buildAbout() termui.Drawable {
	about.SetRect(0, 0, sidebarWidth, 3)

	return about
}
