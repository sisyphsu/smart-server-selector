package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type row struct {
	flex *tview.Flex
	env  *tview.TextView
	host *tview.TextView
	desc *tview.TextView
}

func newRow() *row {
	r := &row{
		flex: tview.NewFlex(),
		env:  tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorGreen),
		host: tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorGreen),
		desc: tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorGreen),
	}
	r.flex.SetDirection(tview.FlexColumn)
	r.flex.AddItem(r.env, 12, 1, false)
	r.flex.AddItem(r.host, 36, 1, false)
	r.flex.AddItem(r.desc, 0, 10, false)

	r.env.SetBackgroundColor(tcell.ColorDefault)
	r.host.SetBackgroundColor(tcell.ColorDefault)
	r.desc.SetBackgroundColor(tcell.ColorDefault)

	return r
}

func (r *row) render(s *server, selected bool, keyword []string) {
	if selected {
		r.flex.SetBackgroundColor(tcell.ColorBlue)
	} else {
		r.flex.SetBackgroundColor(tcell.ColorDefault)
	}
	var env, host, desc string
	if s != nil {
		env = s.env
		host = s.host
		desc = s.desc
		if len(s.user) > 0 {
			host = s.user + "@" + s.host
		}
		if len(s.port) > 0 {
			host = host + ":" + s.port
		}
	}
	r.env.SetText(env)
	r.host.SetText(host)
	r.desc.SetText(desc)
}
