package selector

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
	"sort"
	"strings"
)

type ServersUI struct {
	flex *tview.Flex
	rows []*tview.TextView

	offset  int
	keyword string
	visible []server
	all     []server
}

func newServersUI(all []server) *ServersUI {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true).
		SetBorderColor(tcell.ColorDarkCyan).
		SetBackgroundColor(tcell.ColorBlack)
	v := &ServersUI{flex: flex, all: all}

	app.SetAfterDrawFunc(func(screen tcell.Screen) {
		v.render()
	})

	return v
}

func (s *ServersUI) onEvent(event *tcell.EventKey) bool {
	l := len(s.visible)
	switch event.Key() {
	case tcell.KeyDown, tcell.KeyTab, tcell.KeyPgDn: // select down
		s.offset = (s.offset + 1 + l) % l
		s.render()
		return true
	case tcell.KeyUp, tcell.KeyBacktab, tcell.KeyPgUp: // select up
		s.offset = (s.offset - 1 + l) % l
		s.render()
		return true
	}
	return false
}

func (s *ServersUI) setKeyword(kw string) {
	s.offset = 0
	s.keyword = kw
	s.render()
}

func (s *ServersUI) setServers(all []server) {
	s.offset = 0
	s.all = all
	s.render()
}

func (s *ServersUI) flushVisible() {
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
}

func (s *ServersUI) flushRows() {
	_, _, _, height := s.flex.GetInnerRect()
	for l := len(s.rows); l > height; {
		l--
		s.flex.RemoveItem(s.rows[l])
		s.rows = s.rows[:l]
	}
	for l := len(s.rows); l < height; {
		l++
		newRow := tview.NewTextView().SetDynamicColors(true)
		s.flex.AddItem(newRow, 1, 0, false)
		s.rows = append(s.rows, newRow)
	}
}

func (s *ServersUI) render() {
	s.flushRows()
	s.flushVisible()
	servers := s.visible
	offset := s.offset
	if rowNum := len(s.rows); offset >= rowNum {
		servers = servers[offset-rowNum : offset]
		offset = rowNum - 1
	}
	for i, row := range s.rows {
		if i >= len(servers) {
			row.SetBackgroundColor(tcell.ColorDefault)
			row.SetText("")
			continue
		}
		if i == offset {
			row.SetBackgroundColor(tcell.ColorBlue)
		} else {
			row.SetBackgroundColor(tcell.ColorDefault)
		}
		s := servers[i]
		env := forceWidth(s.env, 6)
		host := forceWidth(s.host, 16)
		user := forceWidth(s.user, 10)
		port := forceWidth(s.port, 6)
		desc := s.desc
		row.SetText(fmt.Sprintf(" [yellow]> [green]%v %v %v %v %v", env, host, user, port, desc))
	}
}

func forceWidth(s string, w int) string {
	if l := len(s); l <= w {
		return runewidth.FillRight(s, w)
	} else {
		return s
	}
}
