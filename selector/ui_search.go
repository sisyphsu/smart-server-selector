package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func buildSearchUI() tview.Primitive {
	input := tview.NewInputField().
		SetPlaceholder("env, host, desc...").
		SetChangedFunc(view.setKeyword).SetFieldBackgroundColor(tcell.ColorBlack)
	input.SetBorder(true).
		SetBorderAttributes(tcell.AttrNone).
		SetBorderColor(tcell.ColorDarkCyan)
	input.SetTitle("Search").
		SetTitleColor(tcell.ColorDarkCyan).
		SetTitleAlign(tview.AlignLeft)

	return input
}
