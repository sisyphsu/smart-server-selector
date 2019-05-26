package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/sisyphsu/server-selector/core"
)

func main() {
	// init ui
	if err := ui.Init(); err != nil {
		println("failed to initialize termui: ", err)
		return
	}
	defer ui.Close()

	// init render
	core.Render()

	// event loop
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		exit, update := core.HandleEvent(e)
		if exit {
			break
		}
		if update {
			core.Render()
		}
	}
}
