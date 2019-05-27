package selector

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"sort"
	"strings"
)

type server struct {
	env   string
	host  string
	desc  string
	score int
}

type ServerTable struct {
	selected int
	keyword  string
	visible  []server
	all      []server

	table *widgets.Table
}

func NewServerTable(servers []server) *ServerTable {
	st := new(ServerTable)
	st.table = widgets.NewTable()
	st.table.RowSeparator = false
	st.table.Title = "Servers"
	st.table.TitleStyle.Fg = ui.ColorCyan
	st.table.TextAlignment = ui.AlignLeft
	st.table.Border = true
	st.table.BorderStyle.Fg = ui.ColorCyan

	st.all = servers
	st.flushVisible()

	return st
}

func (st *ServerTable) setKeyword(kw string) {
	st.keyword = kw
	st.flushVisible()
	st.render()
}

func (st *ServerTable) setServers(ss []server) {
	st.all = ss
	st.flushVisible()
	st.render()
}

func (st *ServerTable) onEvent(event ui.Event) {
	if !Front {
		return
	}
	switch event.ID {
	case "<Down>", "<Tab>": // select down
		st.moveSelect(1)
	case "<Up>": // select up
		st.moveSelect(-1)
	case "<Enter>": // confirm, ok

	}
}

func (st *ServerTable) moveSelect(off int) {
	l := len(st.visible)
	st.selected = (st.selected + off + l) % l
	st.render()
}

func (st *ServerTable) flushVisible() {
	kws := strings.Split(strings.TrimSpace(st.keyword), " ")
	for i, kw := range kws {
		kws[i] = strings.TrimSpace(kw)
	}
	var result []server
	for _, server := range st.all {
		server.score = 0
		for _, kw := range kws {
			if strings.Contains(server.env, kw) {
				server.score += 1000
			}
			if strings.Contains(server.host, kw) {
				server.score += 300
			}
			if strings.Contains(server.desc, kw) {
				server.score += 100
			}
		}
		if server.score > 0 || len(kws) == 0 {
			result = append(result, server)
		}
	}
	sort.Sort(servers(result))

	st.selected = 0
	st.visible = result
}

func (st *ServerTable) render() {
	if !Front {
		return
	}
	var data [][]string
	for i, server := range st.visible {
		data = append(data, []string{server.env, server.host, server.desc})
		if st.selected == i {
			st.table.RowStyles[i] = ui.Style{Bg: ui.ColorRed, Fg: ui.ColorWhite, Modifier: ui.ModifierBold}
		} else {
			st.table.RowStyles[i] = ui.Style{Bg: ui.ColorClear, Fg: ui.ColorGreen, Modifier: ui.ModifierClear}
		}
	}
	st.table.SetRect(sidebarWidth, 3, termWidth(), termHeight())
	st.table.Rows = data

	ui.Render(st.table)
}
