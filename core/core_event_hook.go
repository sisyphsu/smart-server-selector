package core

import "github.com/gizak/termui"

func init() {
	// page down
	termui.Handle("/sys/kbd/<down>", func(event termui.Event) {
		println("page down")
	})
	// page up
	termui.Handle("/sys/kbd/<up>", func(event termui.Event) {
		println("page up")
	})

	// enter
	termui.Handle("/sys/kbd/<enter>", func(event termui.Event) {
		println("enter")
	})
}

// Handle all events
func EventHandler(e termui.Event) {
	// page down
	switch e.Path {
	case "/sys/kbd/<down>":
		// select down
	case "/sys/kbd/<up>":
		// select up
	case "/sys/kbd/<left>":
		// select ring left, if has left
	case "/sys/kbd/<right>":
		// select ring next, if has next
	case "/sys/kbd/<enter>":
		// confirm, ok
	case "/sys/kbd/<space>":
		// keyword + <space>
	case "/sys/kbd/<tab>":
		// select next option
	case "/sys/kbd/<escape>":
		// clear keyword
	case "/sys/kbd/C-8", "/sys/kbd/<delete>":

	default:

	}
	// 2. start server list editor
	// 3. keyword input
	// 4. up/down select
	println(e)
}
