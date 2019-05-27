package selector

import (
	ui "github.com/gizak/termui/v3"
)

var hintsStr = []string{
	"  [Ctrl+C] exit",
	"  [Ctrl+P] edit",
	"      [Up] prev",
	"    [Down] next",
	"     [Tab] switch",
	"     [Esc] clear",
	"     [Del] delete",
	"   [Enter] confirm",
	"      [....](fg-white)    ",
}

var testServers = []server{
	{env: "test", host: "172.10.10.130", desc: "admin,Front,user"},
	{env: "test", host: "172.10.10.131", desc: "trade,user,mysql"},
	{env: "test", host: "172.10.10.132", desc: "trade,redis,zookeeper"},
	{env: "pre", host: "172.10.40.45", desc: "admin"},
	{env: "pre", host: "172.10.40.46", desc: "user, trade"},
	{env: "pre", host: "172.10.40.47", desc: "trade, search"},
	{env: "prod", host: "172.10.40.203", desc: "admin"},
	{env: "prod", host: "172.10.40.204", desc: "user"},
	{env: "prod", host: "172.10.40.205", desc: "user, search"},
	{env: "prod", host: "172.10.40.206", desc: "trade, search"},
	{env: "prod", host: "172.10.40.207", desc: "trade"},
}

func loadServers() []server {
	return testServers
}

var EventHandler []func(ui.Event)

func termWidth() int {
	w, _ := ui.TerminalDimensions()
	return w
}

func termHeight() int {
	_, h := ui.TerminalDimensions()
	return h
}

type servers []server

func (a servers) Len() int      { return len(a) }
func (a servers) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a servers) Less(i, j int) bool {
	if a[i].score != a[j].score {
		return a[i].score < a[j].score
	}
	if a[i].env != a[j].env {
		return a[i].env < a[j].env
	}
	if a[i].host != a[j].host {
		return a[i].host < a[j].host
	}
	return i < j
}
