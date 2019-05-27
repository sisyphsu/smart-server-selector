package main

import (
	"github.com/gizak/termui"
	"github.com/sisyphsu/server-selector/selector"
)

func main() {
	// init ui
	if err := termui.Init(); err != nil {
		println("failed to initialize termui: ", err)
		return
	}
	defer termui.Close()

	// init draw
	go selector.Start()

	// loop ui
	termui.Loop()
}
