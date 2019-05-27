package selector

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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

	table *widgets.Table
}

func NewServerTable(servers []server) *ServerTable {
	st := new(ServerTable)
	st.table = widgets.NewTable()
	st.table.RowSeparator = false
	st.table.Title = "Servers"
	st.table.TitleStyle.Fg = ui.ColorCyan
	st.table.TextAlignment = ui.AlignLeft
	st.table.TextStyle.Fg = ui.ColorGreen
	st.table.Border = true
	st.table.BorderStyle.Fg = ui.ColorCyan

	st.serverAll = servers

	//termui.Handle("/sys/kbd", st.onEvent)

	return st
}

func (st *ServerTable) setKeyword(kw string) {
	st.keyword = kw
	st.render()
}

func (st *ServerTable) onEvent(event ui.Event) {
	if !Front {
		return
	}
	switch event.ID {
	case "<Down>": // select down

	case "<Up>": // select up

	case "<Tab>": // select next option

	case "<Enter>": // confirm, ok

	}
}

func (st *ServerTable) render() {
	if !Front {
		return
	}
	var data [][]string
	for _, server := range st.serverAll {
		data = append(data, []string{server.env, server.host, server.desc})
	}
	st.table.SetRect(sidebarWidth, 3, termWidth(), termHeight())
	st.table.Rows = data
	// handle selected
	for i, row := range data {
		if row[1] != st.serverSelected {
			continue
		}
		st.table.RowStyles[i] = ui.Style{Bg: ui.ColorRed, Fg: ui.ColorYellow}
	}
	ui.Render(st.table)
}
