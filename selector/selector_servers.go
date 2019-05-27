package selector

import "github.com/gizak/termui"

type server struct {
	env  string
	host string
	desc string
}

type ServerTable struct {
	keyword        string
	serverSelected string
	serverVisible  []server
	serverAll      []server

	table *termui.Table
}

func NewServerTable(servers []server) *ServerTable {
	st := new(ServerTable)
	st.table = termui.NewTable()
	st.table.Separator = false
	st.table.Border = true
	st.table.BorderLabel = "Servers"
	st.table.BorderLabelFg = termui.ColorCyan
	st.table.FgColor = termui.ColorGreen
	st.table.BgColor = termui.ColorDefault
	st.table.TextAlign = termui.AlignLeft
	st.table.BorderFg = termui.ColorCyan

	st.serverAll = servers

	termui.Handle("/sys/kbd", st.onShortcut)

	return st
}

func (st *ServerTable) setKeyword(kw string) {
	st.keyword = kw
	st.render()
}

func (st *ServerTable) onShortcut(event termui.Event) {
	if !front {
		return
	}
	switch event.Path {
	case "/sys/kbd/<down>": // select down

	case "/sys/kbd/<up>": // select up

	case "/sys/kbd/<enter>": // confirm, ok

	case "/sys/kbd/<tab>": // select next option

	}
}

func (st *ServerTable) render() {
	if !front {
		return
	}
	var data [][]string
	for _, server := range st.serverAll {
		data = append(data, []string{server.env, server.host, server.desc})
	}
	st.table.X = sidebarWidth
	st.table.Y = 3
	st.table.Rows = data
	st.table.Analysis()
	st.table.SetSize()
	st.table.Width = termui.TermWidth() - sidebarWidth
	// handle selected
	for i, row := range data {
		if row[1] != st.serverSelected {
			continue
		}
		st.table.BgColors[i] = termui.ColorRed
		st.table.FgColors[i] = termui.ColorYellow
	}
	termui.Render(st.table)
}
