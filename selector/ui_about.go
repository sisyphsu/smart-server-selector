package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func buildAboutUI() *tview.TextView {
	tx := tview.NewTextView()
	tx.SetTitle("About")
	tx.SetTitleColor(tcell.ColorDarkCyan)
	tx.SetTitleAlign(tview.AlignCenter)
	tx.SetText("Smart Server Selector")
	tx.SetTextAlign(tview.AlignCenter)
	tx.SetTextColor(tcell.ColorYellow)
	tx.SetBorder(true)
	tx.SetBorderColor(tcell.ColorDarkCyan)

	return tx
}
