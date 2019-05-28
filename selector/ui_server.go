package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"sort"
	"strings"
)

type ServerUI struct {
	flex *tview.Flex
	rows []*row

	offset  int
	keyword string
	kws     []string
	visible []server
	all     []server
}

func newServersUI(all []server) *ServerUI {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true).
		SetBorderColor(tcell.ColorDarkCyan).
		SetBackgroundColor(tcell.ColorBlack)
	v := &ServerUI{flex: flex, all: all}

	app.SetAfterDrawFunc(func(screen tcell.Screen) {
		v.render()
	})

	return v
}

func (s *ServerUI) onEvent(event *tcell.EventKey) bool {
	l := len(s.visible)
	switch event.Key() {
	case tcell.KeyDown, tcell.KeyTab, tcell.KeyPgDn: // select down
		s.selectOffset((s.offset + 1 + l) % l)
		return true
	case tcell.KeyUp, tcell.KeyBacktab, tcell.KeyPgUp: // select up
		s.selectOffset((s.offset - 1 + l) % l)
		return true
	}
	return false
}

func (s *ServerUI) setKeyword(kw string) {
	s.keyword = kw
	s.selectOffset(0)
}

func (s *ServerUI) setServers(all []server) {
	s.all = all
	s.selectOffset(0)
}

func (s *ServerUI) selectOffset(off int) {
	s.offset = off
	s.render()
}

func (s *ServerUI) flushVisible() {
	var kws = make([]string, 0)
	for _, kw := range strings.Split(s.keyword, " ") {
		kw = strings.TrimSpace(kw)
		if len(kw) == 0 {
			continue
		}
		kws = append(kws, kw)
	}
	var result []server
	for _, server := range s.all {
		server.score = 0
		for _, kw := range kws {
			kw = strings.ToLower(kw)
			if strings.Contains(strings.ToLower(server.env), kw) {
				server.score += 1000
			}
			if strings.Contains(strings.ToLower(server.host), kw) {
				server.score += 300
			}
			if strings.Contains(strings.ToLower(server.desc), kw) {
				server.score += 100
			}
		}
		if server.score > 0 || len(kws) == 0 {
			result = append(result, server)
		}
	}
	sort.Sort(serverArray(result))

	s.visible = result
	s.kws = kws
}

func (s *ServerUI) flushRows() {
	_, _, _, height := s.flex.GetInnerRect()
	for l := len(s.rows); l > height; {
		l--
		s.flex.RemoveItem(s.rows[l].flex)
		s.rows = s.rows[:l]
	}
	for l := len(s.rows); l < height; {
		l++
		newRow := newRow()
		s.flex.AddItem(newRow.flex, 1, 0, false)
		s.rows = append(s.rows, newRow)
	}
}

func (s *ServerUI) render() {
	s.flushRows()
	s.flushVisible()
	servers := s.visible
	offset := s.offset
	if rowNum := len(s.rows); offset >= rowNum {
		servers = servers[offset-rowNum : offset]
		offset = rowNum - 1
	}
	for i, row := range s.rows {
		var selected = i == offset
		var curr *server
		if i < len(servers) {
			tmp := servers[i]
			curr = &tmp
		}
		row.render(curr, selected, s.kws)
	}
}
