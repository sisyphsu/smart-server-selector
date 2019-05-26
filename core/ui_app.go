package core

import (
	"fmt"
	"github.com/gizak/termui"
)

type Selector struct {
	about       *termui.Par
	hints       *termui.List
	keyword     *Keyword
	serverTable *ServerTable
}

func NewSelector() *Selector {
	if err := termui.Init(); err != nil {
		panic(fmt.Sprintf("failed to initialize termui: %v", err))
	}

	s := new(Selector)
	// about
	s.about = termui.NewPar("Smart Server Selector")
	s.about.X = 0
	s.about.Height = 3
	s.about.Width = sidebarWidth
	s.about.TextFgColor = termui.ColorYellow
	s.about.BorderLabel = "About"
	s.about.BorderLabelFg = termui.ColorCyan
	s.about.BorderFg = termui.ColorCyan
	termui.Render(s.about)
	// hints
	s.hints = termui.NewList()
	s.hints.Items = hintsStr
	s.hints.ItemFgColor = termui.ColorYellow
	s.hints.BorderLabel = "Hints"
	s.hints.BorderLabelFg = termui.ColorCyan
	s.hints.BorderFg = termui.ColorCyan
	s.hints.Y = 3
	s.hints.X = 0
	s.hints.Width = sidebarWidth
	s.hints.Height = termui.TermHeight() - 3
	termui.Render(s.hints)
	// keyword
	s.keyword = NewKeyword()
	// table
	s.serverTable = NewServerTable()

	return s
}

func (s *Selector) Run() {
	termui.DefaultEvtStream.Hook(func(event termui.Event) {
		//switch event.Path {
		//case "/timer/1s":
		//	update = true
		//case "/sys/wnd/resize":
		//	update = true
		//case "/sys/kbd/C-c":
		//	exit = true // control-c force exit
		//}
	})
	termui.Loop()
}

func (s *Selector) Close() {
	s.keyword.close()

	termui.Clear()
	termui.Close()
}
