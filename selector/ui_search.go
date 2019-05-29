package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var searchInput *tview.InputField

func buildSearchUI() *tview.InputField {
	searchInput = tview.NewInputField().
		SetPlaceholder("env, host, desc...").
		SetChangedFunc(view.setKeyword).SetFieldBackgroundColor(tcell.ColorBlack)
	searchInput.SetBorder(true).
		SetBorderAttributes(tcell.AttrNone).
		SetBorderColor(tcell.ColorDarkCyan)
	searchInput.SetTitle("Search").
		SetTitleColor(tcell.ColorDarkCyan).
		SetTitleAlign(tview.AlignLeft)

	return searchInput
}
