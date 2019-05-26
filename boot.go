package main

import (
	"github.com/gizak/termui"
	"github.com/sisyphsu/server-selector/core"
)

func main() {
	// init ui
	if err := termui.Init(); err != nil {
		println("failed to initialize termui: ", err)
		return
	}
	defer termui.Close()

	// init draw
	core.Draw()
	core.NewKeyword()

	// event loop
	termui.DefaultEvtStream.Hook(func(event termui.Event) {
		var exit, update bool
		switch event.Path {
		case "/timer/1s":
			update = true
		case "/sys/wnd/resize":
			update = true
		case "/sys/kbd/C-c":
			exit = true // control-c force exit
		default:
			exit, update = core.HandleEvent(event)
		}
		if exit {
			termui.StopLoop() // close
		}
		if update {
			core.Draw() // render
		}
	})
	termui.Loop()
}
