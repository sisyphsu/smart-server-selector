package selector

import "testing"

func TestHighlight(t *testing.T) {
	s := "server selector is so smart"
	kw := []string{"se", "e", "rt"}
	words := splitKws(s, kw)
	for _, w := range words {
		t.Log(w.highlight, w.txt)
	}
}
