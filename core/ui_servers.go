package core

import "github.com/gizak/termui"

type ServerTable struct {
	keyword        string
	serverSelected string
	serverVisible  []server
	serverAll      []server

	table *termui.Table
}

func NewServerTable() *ServerTable {
	st := new(ServerTable)
	st.table = termui.NewTable()
	st.table.X = sidebarWidth
	st.table.Y = 3
	st.table.Separator = false
	st.table.Border = true
	st.table.BorderLabel = "Servers"
	st.table.BorderLabelFg = termui.ColorCyan
	st.table.FgColor = termui.ColorGreen
	st.table.BgColor = termui.ColorDefault
	st.table.TextAlign = termui.AlignLeft
	st.table.BorderFg = termui.ColorCyan

	st.serverAll = loadServers() // can't change

	return st
}

func (st *ServerTable) SetKeyword(kw string) {
	st.keyword = kw
	st.update()
}

func (st *ServerTable) SetActive(server string) {
	st.serverSelected = server
}

func (st *ServerTable) update() {
	var data [][]string
	for _, server := range st.serverAll {
		data = append(data, []string{server.env, server.host, server.desc})
	}
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
