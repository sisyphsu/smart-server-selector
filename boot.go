package main

import (
	"github.com/gizak/termui"
	"github.com/sisyphsu/server-selector/core"
)

func main() {
	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()
	// init core

	// build layout
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(6, 0, nil),
			termui.NewCol(6, 0, nil)),
		termui.NewRow(
			termui.NewCol(3, 0, nil),
			termui.NewCol(3, 0, nil),
			termui.NewCol(6, 0, nil)))

	// hook event
	termui.DefaultEvtStream.Hook(core.EventHandler)

	// loop
	termui.Loop()
}
