package core

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
	{"test", "172.10.10.130", "admin,front,user"},
	{"test", "172.10.10.131", "trade,user,mysql"},
	{"test", "172.10.10.132", "trade,redis,zookeeper"},
	{"pre", "172.10.40.45", "admin"},
	{"pre", "172.10.40.46", "user, trade"},
	{"pre", "172.10.40.47", "trade, search"},
	{"prod", "172.10.40.203", "admin"},
	{"prod", "172.10.40.204", "user"},
	{"prod", "172.10.40.205", "user, search"},
	{"prod", "172.10.40.206", "trade, search"},
	{"prod", "172.10.40.207", "trade"},
}
