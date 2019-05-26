package main

import (
	"github.com/gizak/termui"
	"github.com/sisyphsu/server-selector/core"
)

func main() {
	// init ui
	defer termui.Close()

	// init draw
	for {
		selector := core.NewSelector()
		selector.Run()
	}
}
