package selector

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var tips = [][]string{
	{"Ctrl+P", "edit servers"},
	{"Up", "switch previous"},
	{"Down", "switch next"},
	{"Shift+Tab", "switch previous"},
	{"Tab", "switch to next"},
	{"Enter", "confirm select"},
	{"Ctrl+C", "exit"},
	{"Esc+Esc", "exit"},
}

func buildTipsUI() tview.Primitive {
	l := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	l.SetTitle("Tips").
		SetTitleAlign(tview.AlignCenter).
		SetTitleColor(tcell.ColorDarkCyan)
	l.SetBorder(true).
		SetBorderColor(tcell.ColorDarkCyan)

	for i, tip := range tips {
		_, _ = fmt.Fprintf(l, " [yellow]%v. [white]%v\n", i+1, tip[0])
		_, _ = fmt.Fprintf(l, "    [green]%v\n", tip[1])
	}
	return l
}
