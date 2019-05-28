package selector

type server struct {
	env   string
	host  string
	port  string
	user  string
	desc  string
	score int
}

type serverArray []server

func (a serverArray) Len() int      { return len(a) }
func (a serverArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a serverArray) Less(i, j int) bool {
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
