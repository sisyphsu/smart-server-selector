package main

import (
	"github.com/gizak/termui"
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	rows2 := [][]string{
		{"header1", "header2", "header3"},
		{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		{"2016", "11", "11"},
	}

	table2 := termui.NewTable()
	table2.Rows = rows2
	table2.FgColor = termui.ColorGreen
	table2.BgColor = termui.ColorDefault
	table2.TextAlign = termui.AlignLeft
	table2.Separator = false
	table2.Analysis()
	table2.SetSize()
	table2.BgColors[2] = termui.ColorRed
	table2.Y = 10
	table2.X = 0
	table2.Border = true
	table2.BorderLabel = "servers"
	table2.Width = termui.TermWidth()

	offset := 0
	termui.Render(table2)
	termui.DefaultEvtStream.Hook(func(event termui.Event) {
		if event.Path == "/sys/kbd/q" {
			termui.StopLoop()
		} else if event.Path == "/sys/kbd/<down>" {
			offset = (offset + 1) % len(table2.Rows)
			for i := range table2.BgColors {
				if offset == i {
					table2.BgColors[i] = termui.ColorRed
				} else {
					table2.BgColors[i] = termui.ColorDefault
				}
			}
			termui.Clear()
			termui.Render(table2)
		} else if event.Path == "/sys/wnd/resize" {
			table2.Width = termui.TermWidth()

			termui.Clear()
			termui.Render(table2)
		} else {
			println(event.Path)
		}
	})

	termui.Loop()
}
