package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
	"github.com/sisyphsu/smart-server-selector/selector"
)

func main() {
	runewidth.DefaultCondition.EastAsianWidth = false
	app := tview.NewApplication()

	selector.Start(app)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
