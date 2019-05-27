package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/sisyphsu/server-selector/selector"
)

func main() {
	// init ui
	if err := ui.Init(); err != nil {
		println("failed to initialize termui: ", err)
	}
	defer ui.Close()

	// init draw
	selector.Start()
}
