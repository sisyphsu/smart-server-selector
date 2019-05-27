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

	return st
}

func (st *ServerTable) setKeyword(kw string) {
	st.keyword = kw
	st.render()
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
		if row[1] != cxt.serverSelected {
			continue
		}
		st.table.BgColors[i] = termui.ColorRed
		st.table.FgColors[i] = termui.ColorYellow
	}
	termui.Render(st.table)
}
