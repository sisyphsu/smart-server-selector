package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strings"
)

// an word split from str
type word struct {
	txt       string
	highlight bool
}

// row represent an server in server list, it wrapped one server's render logic.
type row struct {
	flex *tview.Flex
	env  *tview.TextView
	host *tview.TextView
	desc *tview.TextView
}

// create an new row
func newRow() *row {
	r := &row{
		flex: tview.NewFlex(),
		env:  tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorLawnGreen),
		host: tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorLawnGreen),
		desc: tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorLawnGreen),
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

// render the current row
func (r *row) render(s *server, selected bool, kws []string) {
	if selected {
		r.flex.SetBackgroundColor(tcell.ColorRoyalBlue)
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
	r.env.SetText(highlight(env, kws))
	r.host.SetText(highlight(host, kws))
	r.desc.SetText(highlight(desc, kws))
}

// generate highlight text for the specified string
func highlight(s string, kws []string) (r string) {
	for _, word := range splitKws(s, kws) {
		if word.highlight {
			r += "[red]"
		} else {
			r += "[lawngreen]"
		}
		r += word.txt
	}
	return
}

// split the specified string with kws
func splitKws(s string, kws []string) []word {
	result := []word{{txt: s}}
	for _, kw := range kws {
		tmp := make([]word, 0)
		for _, w := range result {
			if w.highlight {
				tmp = append(tmp, w)
				continue
			}
			parts := strings.Split(w.txt, kw)
			for i, part := range parts {
				if i > 0 {
					tmp = append(tmp, word{txt: kw, highlight: true})
				}
				if len(part) > 0 {
					tmp = append(tmp, word{txt: part, highlight: false})
				}
			}
		}
		result = tmp
	}
	return result
}
