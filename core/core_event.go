package core

import (
	"github.com/gizak/termui"
	"strings"
)

// Handle all events
func HandleEvent(event termui.Event) (exit bool, update bool) {
	switch event.Path {
	case "/sys/kbd/<down>": // select down

	case "/sys/kbd/<up>": // select up

	case "/sys/kbd/<left>": // select ring left, if has left

	case "/sys/kbd/<right>": // select ring next, if has next

	case "/sys/kbd/<enter>": // confirm, ok

	case "/sys/kbd/<space>": // keyword + <space>

	case "/sys/kbd/<tab>": // select next option

	case "/sys/kbd/<escape>": // clear keyword
		if l := len(cxt.keyword); l > 0 {
			cxt.keyword = ""
			update = true // update
		}
	case "/sys/kbd/<backspace>", "/sys/kbd/<delete>", "/sys/kbd/C-8": // delete one keyword
		if l := len(cxt.keyword); l > 0 {
			cxt.keyword = cxt.keyword[:len(cxt.keyword)-1]
			update = true // update
		}
	default:
		parts := strings.Split(event.Path, "/")
		if l := len(parts); l > 0 && len(parts[l-1]) == 1 {
			cxt.keyword += parts[l-1]
			update = true // update
		}
	}
	// 2. start server list editor
	// 3. keyword input
	// 4. up/down select
	//println(event.Path)
	return
}
